package autoapi

import(
    "net/http"
)

// RegistryOpts
type RegistryOpts struct {
    // Mux
    Mux *http.ServeMux
    
    // Prefix
    Prefix string
}