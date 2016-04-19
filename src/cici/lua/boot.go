package lua

import (
	"io/ioutil"
	"log"
	"runtime"
	"time"

	gopherlua "github.com/yuin/gopher-lua"
)

// Runs a lua script
func Run(script string) {
	ls := gopherlua.NewState()
	defer ls.Close()

	setGlobals(ls)

	scriptbytes, err := ioutil.ReadFile(script)
	if err != nil {
		log.Fatal("Invalid Script:", err)
	}

	if err := ls.DoString(string(scriptbytes)); err != nil {
		log.Fatal(err)
	}
}

// setGlobals sets top level globals
func setGlobals(ls *gopherlua.LState) {
	ls.SetGlobal("sleep", ls.NewFunction(lsleep))

	ls.SetGlobal("sched", ls.NewFunction(lsched))

	ls.SetGlobal("thread", ls.NewFunction(lthread))

	ls.SetGlobal("print", ls.NewFunction(func(ls *gopherlua.LState) int {
		log.Println(ls.ToString(1))
		return 0
	}))
}

// lsched invokes the golang goroutine scheduler
func lsched(ls *gopherlua.LState) int {
	runtime.Gosched()
	return 0
}

// lsleep
func lsleep(ls *gopherlua.LState) int {
	time.Sleep(time.Duration(ls.ToInt(-1)) * time.Millisecond)
	return 0
}

// lthread calls a lua function within a new lua state which can run against
// the golang goroutine scheduler
func lthread(ls *gopherlua.LState) int {
	thread := gopherlua.NewState()
	ls.XMoveTo(thread, 1)
	threadfunc := thread.ToFunction(-1)

	go func(thread *gopherlua.LState, threadfunc *gopherlua.LFunction) {
		thread.SetGlobal("threadfunc", threadfunc)
		thread.DoString(`threadfunc()`)
		thread.Close()
	}(thread, threadfunc)

	return 0
}
