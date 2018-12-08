package main

import (
  "fmt"
  "net/http"
)

func main() {
  // Processing dynamic requests...
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "This is a simple Http Server application.")
    // can read GET parameters with "r.URL.Querry().Get("token")" 
    // POST parameters(fields from a HTML form)..."f.FromValue("email")"
  })
  // Serving static assets...
  file_server := http.FileServer(http.Dir("static/"))
  // Once our file server is in place, we just need to point a url path at it...
  // this is the name of the directory our files live in.
  http.Handle("static/", http.StripPrefix("/static/", file_server))
  // Accept connections by listening on a port...
  http.ListenAndServe(":80", nil)
}
