package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// logging logs all requests with its path and the time it took to process
func Logging() Middleware {
  // Create a new Middleware
  return func(f http.HandleFunc) http.HandlerFunc {
    // Define the http.HandlerFunc
    return func(w http.ResponseWriter, r *http.Request) {
      start := time.Now()
      defer func() { log.Println(r.URL.Path, time.Since(start)) }
      //call the next middleware/handler in chain
      f(w,r)
    }
  }
}

// Method ensures that url can only be requested with a specific method,
// else returns a 400 Bad Request
func Method(m string) Middleware {
  // create a new middleware...
  return func(f http.HandleFunc) http.HandlerFunc {
    // Define the http.HandlerFunc
    return func(w http.ResponseWriter, r *http.Request) {
      if r.Method != m {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
      }
      f(w, r)
    }
  }
}

func Chain(f http.HandlerFunc, middlewares...Middleware) http.HandlerFunc {
  for_, m := range middlewares {
    f = m(f)
  }
  return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func main() {
  http.HandlerFunc("/", Chain(Hello, Mehtod("GET"), Logging()))
  http.ListenAndServe(":8080", nil)
}
