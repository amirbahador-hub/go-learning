package main

import (
  "fmt"
  "net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func bookhandler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "GodBay World, %s!", request.URL.Path[1:])
}


func main(){
  http.HandleFunc("/", handler)
  http.HandleFunc("/book/", bookhandler)
  http.ListenAndServe(":8080", nil)
}
