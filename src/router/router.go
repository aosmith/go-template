package router

import (
  "fmt"
  "net/http"
)

func ApplicationRouter(writer http.ResponseWriter, request *http.Request) {
  log_request(request)
}

func log_request(r *http.Request) {
  fmt.Printf("Request: %s [%s]\n", r.URL, r.Method)
  fmt.Printf("  Data: %s\n", request_data(r))
  fmt.Printf("Routing request: %s\n", r.URL.Path)
}
