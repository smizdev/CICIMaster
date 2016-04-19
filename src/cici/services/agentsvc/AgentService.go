package agentsvc

import(
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

// AgentService
type AgentService struct {
    // address
    address string
    
    // mux
    mux *http.ServeMux
    
    // wg
    wg *sync.WaitGroup
    
    // opts
    opts *AgentServiceOpts
}

func NewAgentService(opts *AgentServiceOpts) *AgentService {
    return &AgentService{
        address: fmt.Sprintf("%s:%s", opts.BindAddress, strconv.Itoa(opts.BindPort)),
        mux: http.NewServeMux(),
        opts: opts,
    }
}

// main
func (t *AgentService) main() {
    defer func() {
       t.opts.WaitGroup.Done() 
    }()
    
    log.Printf("Agent Service Starting %s:%s\n", t.opts.BindAddress, strconv.Itoa(t.opts.BindPort))
    
    log.Fatal(http.ListenAndServe(t.address, t.mux))
}

// Start
func (t *AgentService) Start() {
    t.opts.WaitGroup.Add(1)
    go t.main()    
}