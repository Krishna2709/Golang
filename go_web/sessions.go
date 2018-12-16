package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/sessions"
)

var (
  // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
  key = []byte("super-secret-key")
  store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "cookie-name")
  // Checking user authentication...
  if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
    http.Error(w, "Forbidden", http.StatusForbidden)
    return
  }
  // Printing the message...
  fmt.Fprintln(w," Welcome to the #100DaysOfCode...Challenge")
}

func login(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "cookie-name")

  session.Values["authenticated"] = true
  session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
  session,_ := store.Get(r, "cookie-name")
  // Revoke the users authentication
  session.Values["authenticated"] = false
  session.Save(r,w)
}

func main() {
  http.HandleFunc("/secret", secret)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)

  http.ListenAndServe(":8080", nil)
}
