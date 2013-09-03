package router

import (
  "fmt"
  "net/http"
)



func main() {
  fmt.Println("Booting router...")
}


func ApplicationRouter(writer http.ResponseWriter, request *http.Request) {
  fmt.Println("Booting router...")
}
