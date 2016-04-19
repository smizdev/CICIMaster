package agentsvc

import(
    "cici/sync"
)

// AgentServiceOpts options for AgentService
type AgentServiceOpts struct {
    // BindAddress
    BindAddress string

    // BindPort
    BindPort int

    // WaitGroup
    WaitGroup sync.WaitGroup
}
