package main

import (
    "fmt"
    "net/http"
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second)://wait for 10 seconds
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}


// What is Context?

// A signal system

// Used to:

// Cancel work if client disconnects

// Enforce timeouts

// Pass request-scoped values

// Here:

// ctx will be cancelled automatically if:

// Client closes browser

// Network drops

// Server times out

