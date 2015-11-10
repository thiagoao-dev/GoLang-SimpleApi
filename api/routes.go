package api

import (
  "fmt"
  "log"
  "time"
  "net/http"
  
  "github.com/gorilla/mux"
)

func RoutesUp() {
  log.Println("Starting routes")
  
  r := mux.NewRouter()
  
  r.HandleFunc("/api", Hello).Methods("OPTIONS")
  UserHandler(r)
  
  http.Handle("/", r)
  
  s := &http.Server {
    Addr: ":8080",
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }
  s.ListenAndServe()
}

func UserHandler(r *mux.Router) {
  r.HandleFunc("/api/user", UserList).Methods("GET")
  r.HandleFunc("/api/user", UserCreate).Methods("POST")
  r.HandleFunc("/api/user/{id:[0-9]+}", UserView).Methods("GET")
  r.HandleFunc("/api/user/{id:[0-9]+}", UserUpdate).Methods("PUT")
  r.HandleFunc("/api/user/{id:[0-9]+}", UserDelete).Methods("DELETE")
}

func Hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World Api")
}

func UserList(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "User List")
}

func UserView(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "User View")
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "User Create")
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "User Update")
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "User Delete")
}