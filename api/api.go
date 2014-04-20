package api

import (
    // "github.com/mrcrilly/fizzbuzz-api-go/fizzbuzz"
    "net/http"
    "log"
    // "encoding/json"
    "regexp"
)

var rx_index    = regexp.MustCompile(`^/$`)
var rx_solo     = regexp.MustCompile(`^/(?P<solo>[0-9]+)[/]?$`)
var rx_range    = regexp.MustCompile(`^/(?P<start>[0-9]+)/(?P<finish>[0-9]+)[/]?$`)

func masterHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the incoming URL path and match against
    // a known URI or return a negative 
    
    if rx_index.MatchString(r.URL.Path) {
        log.Print("Matched the index URL")
        return 
    }

    if rx_solo.MatchString(r.URL.Path) {
        log.Print("Matched the solo URL")
        // solo(w, &r)
        return 
    }

    if rx_range.MatchString(r.URL.Path) {
        log.Print("Matched range URL")
        return 
    }

    log.Print("Matched nothing...")
}

func initHandlers() bool {
    // Using a function for this to handle expansion
    // later on... 

    log.Print("Setting up API handlers")

    http.HandleFunc("/", masterHandler)

    return true
}

func Api() {
    log.Print("Starting API")
    
    initHandlers()
    http.ListenAndServe(":8080", nil)
}