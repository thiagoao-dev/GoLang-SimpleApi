package api

import (
  "log"
  
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

func GetConn() gorm.DB {
  log.Println("Starting database connection")
  
  db, err := gorm.Open("mysql", "u701659451_go:vt1g1dr1@tcp(mysql.hostkeeda.com:3306)/u701659451_go?charset=utf8"); if err != nil {
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