package api

import (
    "net/http"
    "log"
    "regexp"
    "errors"
)

// Used internally to define a URL
// handler resource. We have our compiled
// regex and a function field to the function
// the user wants to call
type URLHandler struct {
    regx *regexp.Regexp 
    handler func(a *Request)
}

// Global variable of our URLHandlers
var Handlers []URLHandler

// Adds a resource handler to the Handlers slice
func AddResource(rx string, h func(a *Request)) (int, error) {
    compiled := regexp.MustCompile(rx)

    i := URLHandler{compiled, h}
    Handlers = append(Handlers, i)

    return 0, nil
}

// Handle the incoming URL path and match against
// a known URI or return a negative 
func masterHandler(w http.ResponseWriter, r *http.Request) {    
    if len(Handlers) >= 1 {
        for _, h := range Handlers {
            m := h.regx.FindStringSubmatch(r.URL.Path)
            if m != nil {
                log.Print("We have matched: ", h.regx.String())
                
                req := &Request{w, r, m}
                h.handler(req)

                return
            }
        }

        log.Print("No matches found.")
    } else {
        log.Print("I have no handlers. Sad face.")
    }
}

func Api() (int, error) {
    log.Print("Starting API")

    // We need some handlers before we can do anything
    if len(Handlers) <= 0 {
        return -1, errors.New("I need some handlers, first")
    }

    // Define the master handler and start listening
    http.HandleFunc("/", masterHandler)
    http.ListenAndServe(":8080", nil)

    return 0, nil
}
