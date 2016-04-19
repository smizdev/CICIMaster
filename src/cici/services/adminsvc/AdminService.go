package adminsvc

import (
	"cici/autoapi"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// AdminService
type AdminService struct {
	// address
	address string

	// opts
	opts *AdminServiceOpts

	// autoApiRegistry
	apireg *autoapi.Registry

	// mux
	mux *http.ServeMux
}

// NewAdminService factory for AdminService
func NewAdminService(opts *AdminServiceOpts) *AdminService {
	return &AdminService{
		address: fmt.Sprintf("%s:%s", opts.BindAddress, strconv.Itoa(opts.BindPort)),
		mux:     http.NewServeMux(),
		opts:    opts,
	}
}

// main
func (t *AdminService) main() {
	defer func() {
		t.opts.WaitGroup.Done()
	}()

	log.Printf("Admin Service Starting %s:%s\n", t.opts.BindAddress, strconv.Itoa(t.opts.BindPort))

	log.Fatal(http.ListenAndServe(t.address, t.mux))
}

// getAgents
func (t *AdminService) getAgents(res http.ResponseWriter, req *http.Request) {

}

// Start starts the admin service
func (t *AdminService) Start() {
	t.opts.WaitGroup.Add(1)
	go t.main()
}
