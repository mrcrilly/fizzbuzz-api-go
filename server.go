package main

import (
    "github.com/mrcrilly/fizzbuzz-api-go/api"
    "github.com/mrcrilly/fizzbuzz-api-go/fizzbuzz"
    "encoding/json"
    "log"
    "strconv"
    "fmt"
)

func indexHandler(request *api.Request) {
    log.Print("OK")
}

func soloHandler(request *api.Request) {
    digit, _ := strconv.Atoi(request.Matches()[1])
    
    if digit <= 0 {
        digit = 1
    }

    results := []string{fizzbuzz.FizzBuzz(digit)}
    passback, err := json.Marshal(results)

    if err != nil {
        return
    }

    // This needs to be handled by the API lib
    request.Out().Header().Set("Content-Type", "application/json")
    fmt.Fprintf(request.Out(), string(passback))
}

func rangeHandler(request *api.Request) {
    start, _ := strconv.Atoi(request.Matches()[1])
    finish, _ := strconv.Atoi(request.Matches()[2])

    if start <= 0 {
        start = 1
    }

    if finish >= 1000 || finish <= 0 {
        finish = 100
    }

    results := make([]string, 0)
    for i := start; i <= finish; i++ {
        results = append(results, fizzbuzz.FizzBuzz(i))
    }

    passback, err := json.Marshal(results)
    
    if err != nil {
        return
    }

    log.Print(string(passback))
}

func main() {
    log.Print("Initialising API")

    api.AddResource(`^/$`, indexHandler)
    api.AddResource(`^/(?P<solo>[0-9]+)[/]?$`, soloHandler)
    api.AddResource(`^/(?P<start>[0-9]+)/(?P<finish>[0-9]+)[/]?$`, rangeHandler)
    api.Api()
}
