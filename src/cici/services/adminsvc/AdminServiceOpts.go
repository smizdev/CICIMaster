package adminsvc

import (
	"cici/sync"
)

// AdminServiceOpts options for AdminService
type AdminServiceOpts struct {
	// BindAddress
	BindAddress string

	// BindPort
	BindPort int

	// WaitGroup
	WaitGroup sync.WaitGroup
}
