package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, you've requested: %s\n Go server is working! \n", r.URL.Path)
}

func HttpDemo() {
    http.HandleFunc("/", helloHandler) // Route
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)  // Start server
}