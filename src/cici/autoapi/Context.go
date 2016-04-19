package autoapi

import(
    "net/http"
)

// Context
type Context struct {
    // Request
   Request *http.Request
   
   // ResponseWriter
   Response http.ResponseWriter 
}