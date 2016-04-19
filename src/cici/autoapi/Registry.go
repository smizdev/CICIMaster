package autoapi

import(
    "sync"
    "fmt"
    "strings"    
    "log"
    "reflect"
)

// Registry
type Registry struct {
    // mutex
    mutex *sync.Mutex
    
    // opts
    opts *RegistryOpts
}

// NewRegistry
func NewRegistry(opts *RegistryOpts) *Registry {
    return &Registry{
        mutex: &sync.Mutex{},
        opts: opts,
    }
}

// Register
func (t *Registry) Register(m interface{}) {
    t.mutex.Lock()
    defer t.mutex.Unlock()
    
    mtype := reflect.TypeOf(m)
    mname := mtype.Elem().Name()
    mpattern := fmt.Sprintf("%s%s", t.opts.Prefix, strings.ToLower(mname))
    
    log.Println("AutoApi Model Registration:", mname , mpattern)
    
    t.opts.Mux.Handle(mpattern, &modelHandler{
        m      : m,        
        mtype  : mtype,
        mget   : mtype.Implements(reflect.TypeOf((*ModelGet)(nil)).Elem()),
        mput   : mtype.Implements(reflect.TypeOf((*ModelPut)(nil)).Elem()),
        mpost  : mtype.Implements(reflect.TypeOf((*ModelPost)(nil)).Elem()),
        mdelete: mtype.Implements(reflect.TypeOf((*ModelDelete)(nil)).Elem()),
        mlist  : mtype.Implements(reflect.TypeOf((*ModelList)(nil)).Elem()),
    })
}

