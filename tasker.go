package main

import (
    "fmt"
    "net/http"
)

func taskerHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a task.")
}

func main() {
    http.HandleFunc("/tasker/", taskerHandler)
    http.ListenAndServe(":8080", nil)
}
