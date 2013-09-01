package main

import (
  "fmt"
  "net/http"
  "handlers"
 // "launchpad.net/goyaml"
)



func main() {
  fmt.Println("Booting server...")
  http.HandleFunc("/", handlers.IndexHandler)
  http.ListenAndServe(":8080", nil)
}
