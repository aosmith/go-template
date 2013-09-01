package main

import (
  "fmt"
  "net/http"
  "router"
 // "launchpad.net/goyaml"
)



func main() {
  fmt.Println("Booting server...")
  http.HandleFunc("/", router.ApplicationRouter)
  http.ListenAndServe(":8080", nil)
}
