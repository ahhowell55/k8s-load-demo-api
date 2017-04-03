package main

import (
  "encoding/json"
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "log"
  "net/http"
)

type Password struct {
  Password       string `json:"plain"`
  HashedPassword string `json:"hashed"`
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    password := r.URL.Path[1:]

    // // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost/2)
    if err != nil {
      panic(err)
    }
    data, _ := json.Marshal(&Password{Password: password, HashedPassword: string(hashedPassword)})
    fmt.Fprintf(w, "%v", string(data))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))

}
