// An example of how to hash & salt your passwords using the bcrypt package in Go.
// This can be used as a login verification system.

package main

import (
  "fmt"
  "log"
  "golang.org/x/crypto/bcrypt" // can get by running in your terminal  -- > go get -u golang.org/x/crypto/bcrypt
)
// Getting the password from the user
func getPassword() []byte {
  fmt.Println("Enter a Password :: ")
  // storing the user input
  var password string

  _, err := fmt.Scan(&password)
  if err!= nil {
    log.Println(err)
  }
// Return the users input as a byte slice
  return []byte(password)
}

//  returns the bcrypt hash of the password at the given cost.
// If the cost given is less than MinCost, the cost will be set to DefaultCost, instead.
// It generates a salt for us which saves us from having to write our own function to generate a salt.
func hashAndSalt(password []byte) string {
  hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
  if err!= nil {
    log.Println(err)
  }
  return string(hash)
}

// CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure
func comparePasswords(hashedPassword string, plainPassword []byte) bool {
// Since we'll be getting the hashed password from the DB it
// will be a string so we'll need to convert it to a byte slice
  byteHash := []byte(hashedPassword)

  err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
  if err != nil {
    log.Println(err)
    return false
  }
  return true
}

func main() {
  // Infinite loop ...it is the same as while (true) {....}
  for {
  pwd := getPassword()
  hash := hashAndSalt(pwd)
  fmt.Println("Salted Hash :: ", hash)

  pwd2 := getPassword()

  pwdMatch := comparePasswords(hash, pwd2)

  fmt.Println("Passwords Match ?", pwdMatch)

  }
}
