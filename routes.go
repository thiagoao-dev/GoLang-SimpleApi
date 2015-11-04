package main

import (
  "fmt"
  "time"
  "net/http"
  
  "github.com/gorilla/mux"
)

func UpRoutes() {
  r := mux.NewRouter()
  r.HandleFunc("/api", Hello).Methods("GET")
  //r.HandleFunc("/api/user/create", CreateUser).Methods("GET")
  //r.HandleFunc("/api/user/{id:[0-9]+}", GetUser).Methods("GET")
  http.Handle("/", r)
  
  s := &http.Server {
    Addr: ":8080",
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }
  s.ListenAndServe()
}

func Hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World Api")
}