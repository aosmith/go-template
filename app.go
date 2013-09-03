package main

import (
  "fmt"
  "net/http"
  "router"
  "os"
 // "launchpad.net/goyaml"
)



func main() {
  args := os.Args
  if len(args) > 1 {
    if args[1] == "test" {
      fmt.Println("Running test suite...")
    } else if args[1] == "generate" {
      fmt.Println("Running generator...")
    } else if args[1] == "server" {
      start_server()
    } else {
      fmt.Printf("Unknown argument: %s\n", args[1])
    }
  } else {
    start_server()
  }
}


func start_server() {
  fmt.Println("Booting server...")
  http.HandleFunc("/", router.ApplicationRouter)
  http.ListenAndServe(":8080", nil)
}

func run_generator(args []string) {
  fmt.Println("Running generator...")
}

func run_test_suite() {
  fmt.Println("Running test suite...")
}
