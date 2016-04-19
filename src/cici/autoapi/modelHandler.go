package autoapi

import(
    "net/http"
    "reflect"
)

// modelHandler
type modelHandler struct {
    // model
    m interface{}  
    
    // modelopts
    mopts *ModelOpts
    
    // modeltype
    mtype reflect.Type
    
    // mget
    mget bool
    
    // mput
    mput bool
    
    // mpost
    mpost bool
    
    // mdelete
    mdelete bool
    
    // mlist
    mlist bool
}

// ServeHTTP
func (t *modelHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    // log.Println("Request for", t.mtype.Elem().Name(), "Get", t.mget, "Put", t.mput, "Post", t.mpost, "Delete", t.mdelete, "List", t.mlist)
    
    if req.Method == "GET" && t.mget {
        t.m.(ModelGet).AutoApiGet(&Context{
            Request: req, 
            Response: res,
        })
    }
}