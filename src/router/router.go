package router

import (
  "fmt"
  "net/http"
)


func ApplicationRouter(writer http.ResponseWriter, request *http.Request) {
  fmt.Printf("Routing request: %s\n", request.URL.Path)

}
