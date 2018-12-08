// Routing (using gorilla/mux)

package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  // creating a new request router..
  r := mux.NewRouter()
  // registering a request handler and URL pattern...
  r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    title := vars["title"] // getting the book title
    page := vars["page"] // getting the page

    fmt.Fprintf(w, "You have requested the book: %s on page: %s\n", title, page)
  })

  // setting up the Server route..
 http.ListenAndServe(":80", r)
}
