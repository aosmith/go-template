package router

import (
  "fmt"
  "net/http"
)




func ApplicationRouter(writer http.ResponseWriter, request *http.Request) {
  fmt.Println("Booting router...")
}
