package main

import (
	"github.com/thiagoao/GoLang-SimpleApi/api"
)

func main() {
  db := api.GetConn();
  
  api.MigrationsUp(&db)
  api.RoutesUp()
}
/*

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
  
  w.Header().Set("Pragma","no-cache")
  w.Header().Set("Access-Control-Allow-Methods", "GET")
  w.Header().Set("Content-Type", "application/json")
  
  urlParams := mux.Vars(r)
  id := urlParams["id"]
  ReadUser := User{}
  
  // Created db connection
  db, err := sql.Open("mysql", "user01:sUtJSsPKnBv9PQXD@tcp(104.197.132.233:3306)/user01")
  if err != nil {
    fmt.Println(err)
  }
  // Ensure closure
  defer db.Close()
 
  err = db.QueryRow("SELECT * FROM users WHERE user_id = ?", id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)
  switch {
    case err == sql.ErrNoRows:
    	fmt.Fprintf(w, "No such user")
    case err != nil:
    	log.Fatal(err)
  		fmt.Fprintf(w, "Error")
    default:
    output, _ := json.Marshal(ReadUser)
    fmt.Fprintf(w, string(output))
  }
  
}

// CreateUser 
func CreateUser(w http.ResponseWriter, r *http.Request) {
  
  NewUser := User{}
  NewUser.Name = r.FormValue("user")
  NewUser.Email = r.FormValue("email")
  NewUser.First = r.FormValue("first")
  NewUser.Last = r.FormValue("last")
  
  output, err := json.Marshal(NewUser)
  fmt.Println(string(output))
  
  if err != nil {
    fmt.Println("Something went wrong")
  }
  
  query := "INSERT INTO users SET user_nickname='" + NewUser.Name + 
    "', user_first='" + NewUser.First + "', user_last='" +
    NewUser.Last + "', user_email='" + NewUser.Email + "'"
  
  // Created db connection
  db, err := sql.Open("mysql", "sql593836:cE9*kN6!@tcp(sql5.freesqldatabase.com:3306)/sql593836")
  if err != nil {
    fmt.Println(err)
  }
  // Ensure closure
  defer db.Close()
  
  // Insert the data
  q, err := db.Exec(query)
  if err != nil {
    fmt.Println(err)
    http.Error(w, err.Error(), http.StatusConflict)
    
  } else {    
    // Print the user exit
    id, err := q.LastInsertId()
    if err != nil {
      fmt.Println("We lost your ID!")
    }
    
    // Recovering ID
    NewUser.ID = id
    user, err := json.Marshal(NewUser)
    fmt.Fprintf(w, string(user))
  }  
}

// Hello
func Hello(w http.ResponseWriter, r *http.Request) {
  
  urlParams := mux.Vars(r)
  name := urlParams["user"]
  HelloMessage := "Hello, " + name
  
  message := API{HelloMessage}
  
  output, err := json.Marshal(message)
  
  if err != nil {
    fmt.Println("Something went wrong!")
  }
  
  fmt.Fprintf(w, string(output))
}
*/