package main

import (
  "fmt"
  "net/http"
  "router"
  "config"
  "os"
)

func main() {
  args := os.Args
  if len(args) > 1 {
    if args[1] == "generate" {
      run_generator(args[1:])
    } else {
      fmt.Printf("Unknown argument: %s\n", args[1])
    }
  } else {
    start_server()
  }
}


func start_server() {
  fmt.Printf("Starting server: %s:%d\n", config.Host, config.Port)
  http.HandleFunc("/", router.ApplicationRouter)
  http.ListenAndServe(config.PortString(), nil)
}

func run_generator(args []string) {
  fmt.Println("Running generator...")
}

func run_test_suite() {
  fmt.Println("Running test suite...")
}
