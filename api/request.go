package api

import (
    "net/http"
    // "log"
    // "errors"
)

// The API's wrapper around net/http.ResponseWriter
// and net/http.Request, as well as the matches from
// the regexp
type Request struct {
    writer http.ResponseWriter
    request *http.Request 
    matches []string
}

func (r Request) Out() http.ResponseWriter {
    return r.writer
}

func (r Request) In() *http.Request {
    return r.request 
}

func (r Request) Matches() []string {
    return r.matches 
}