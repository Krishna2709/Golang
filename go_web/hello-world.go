package main

import (
  "fmt"
  "net/http"
)
func main() {
// Creating a Handler -> receives all incoming HTTP connections from browsers,
//HTTP clients/API reqeusts.
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Request contains all the info..URL/header fields.
    fmt.Fprintf(w, "Hello, you requested: %s\n", r.URL.Path)
  })
// An HTTP server has to listen on a port to pass connections on to the request
  http.ListenAndServe(":80", nil)
}
