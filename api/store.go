package api

import (
  "log"
  
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

func GetConn() gorm.DB {
  log.Println("Starting database connection")
  
  db, err := gorm.Open("mysql", "adminvXTTRbg:qUzAQxxJN6yf@tcp(db-startuping.rhcloud.com:3306)/db?charset=utf8"); if err != nil {
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