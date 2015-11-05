package api

import (
  "log"
  
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

func GetConn() gorm.DB {
  log.Println("Starting database connection")
  
  db, err := gorm.Open("mysql", "golang:G0l@ng@tcp(104.197.24.32:3306)/golang?charset=utf8"); if err != nil {
    log.Fatalln(err)
  }
  defer db.Close()
  
  // Then you could invoke `*sql.DB`'s functions with it
  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)

  // Disable table name's pluralization
  // db.SingularTable(true)

  return db
}