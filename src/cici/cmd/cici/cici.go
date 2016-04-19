package main

import (
	"cici/ioc"
	"cici/lua"
	"cici/services"
	"cici/services/adminsvc"
	"cici/services/agentsvc"
	"cici/sync"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// Dependency Injection Container
var di *ioc.Container = ioc.NewContainer()

// main cici entry point
func main() {
	runtime.GOMAXPROCS(100)

	rand.Seed(time.Now().UTC().UnixNano())

	registration()

	adminsvc := di.Singleton("AdminService").(services.Service)
	adminsvc.Start()

	agentsvc := di.Singleton("AgentService").(services.Service)
	agentsvc.Start()

	lua.Run(os.Args[1])

	di.Singleton("ServiceWaitGroup").(sync.WaitGroup).Wait()
}

// registration
func registration() {
	di.Register("ServiceWaitGroup", func() interface{} {
		return sync.NewWaitGroup()
	})

	di.Register("AdminServiceBindAddress", func() interface{} {
		return "0.0.0.0"
	})

	di.Register("AdminServiceBindPort", func() interface{} {
		return 443
	})

	di.Register("AgentServiceBindAddress", func() interface{} {
		return "0.0.0.0"
	})

	di.Register("AgentServiceBindPort", func() interface{} {
		return 2873
	})

	di.Register("AdminService", func() interface{} {
		return adminsvc.NewAdminService(&adminsvc.AdminServiceOpts{
			BindAddress: di.Singleton("AdminServiceBindAddress").(string),
			BindPort:    di.Singleton("AdminServiceBindPort").(int),
			WaitGroup:   di.Singleton("ServiceWaitGroup").(sync.WaitGroup),
		})
	})

	di.Register("AgentService", func() interface{} {
		return agentsvc.NewAgentService(&agentsvc.AgentServiceOpts{
			BindAddress: di.Singleton("AgentServiceBindAddress").(string),
			BindPort:    di.Singleton("AgentServiceBindPort").(int),
			WaitGroup:   di.Singleton("ServiceWaitGroup").(sync.WaitGroup),
		})
	})
}
