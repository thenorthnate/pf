// vault server
// Nathan North

package main

import (
  "log"
  "net/http"
  "regex"
  "encoding/json"
  "github.com/thenorthnate/govault"
)

type ClientRequestWarning struct {
  Warning string
}


var configPath string = "conf.toml"



func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/api/auth/", loginHandler)
  http.HandleFunc("/api/user/", userHandler)
  http.HandleFunc("/api/account/", accountHandler)
  http.HandleFunc("/api/transaction/", transactionHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func authHandler(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
  // First check to make sure the request was made to a valid URL
  matched := govault.ValidatePath(r.URL.Path, "auth")
  if !matched {
    js, err := json.Marshal(ClientRequestWarning{"Invalid URL for authentication"})
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
    return
  }
  username, password, ok := r.BasicAuth()
}

func userHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
