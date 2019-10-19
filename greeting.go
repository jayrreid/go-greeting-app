package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "log"
)

type Greeting struct {
    Greet string
}


func greeting(w http.ResponseWriter, r *http.Request) {
  // return "hello" greeting
  log.Println("greeting service called")
  data := Greeting{"hola"}
  b, err := json.Marshal(data)
  if err != nil {
    http.Error(w, "Internal Server Error", 500)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, string(b))
}


func main() {
  log.Println("greeting service started")
  http.HandleFunc("/greeting", greeting)
  http.ListenAndServe(":8081", nil)
}
