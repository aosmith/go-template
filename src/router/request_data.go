package router

import (
  "fmt"
  "net/http"
)


func request_data(r *http.Request) []Param {
  var params []Param
  if r.Form != nil {
    fmt.Println("Found form data!")
  }
  return params
}
