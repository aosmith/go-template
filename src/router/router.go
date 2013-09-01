package router

import (
  "net/http"
)


func ApplicationRouter(writer http.ResponseWriter, request *http.Request) {
  if request.Method == "GET" {
    // This is a request for a resource
  } else if request.Method == "POST" {
    // This is a request to create a resource
  } else if request.Method == "PUT" {
    // This is a request to update a resource
  } else if request.Method == "DELETE" {
    // This is a request to delete a resource
  }
}
