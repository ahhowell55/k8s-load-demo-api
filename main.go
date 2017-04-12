package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "github.com/juju/loggo"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "golang.org/x/crypto/bcrypt"
  "log"
  "net/http"
)

type Password struct {
  Password       string `json:"plain"`
  HashedPassword string `json:"hashed"`
}

func main() {
  debug := flag.Bool("debug", false, "Are we in debug mode?")
  flag.Parse()

  baseLogLevel := "WARNING"
  if *debug {
    baseLogLevel = "DEBUG"
  }
  loggo.ConfigureLoggers(fmt.Sprintf("<root>=%v", baseLogLevel))

  logger := loggo.GetLogger("http.PasswordHasher")
  http.Handle("/metrics", promhttp.Handler())
  http.Handle("/", prometheus.InstrumentHandlerFunc("PasswordHasher", func(w http.ResponseWriter, r *http.Request) {
    password := r.URL.Path[1:]
    logger.Debugf("Request Received for password '%s'", password)
    // // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost/2)
    if err != nil {
      logger.Errorf("Failed to bcrypt '%s'. Error: %v", password, err.Error())
    }
    data, _ := json.Marshal(&Password{Password: password, HashedPassword: string(hashedPassword)})
    logger.Debugf("Hashed password '%s' into '%s'", password, hashedPassword)
    fmt.Fprintf(w, "%s", string(data))
  }))

  log.Fatal(http.ListenAndServe(":8080", nil))

}
