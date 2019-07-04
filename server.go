package main

import (
  "fmt"
  "net/http"
  "os"
)

var current_dir = "."

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
  fmt.Fprintf(w, "Current directory is %s!", current_dir)
}

func main() {
  pwd, err := os.Getwd()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  //fmt.Println(pwd)
  current_dir = pwd

  http.HandleFunc("/", handler)
  fmt.Println("Server running...")
  http.ListenAndServe(":8080", nil)
}
