package api

import (
  "log"
  
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

func GetConn() gorm.DB {
  log.Println("Starting database connection")
  
  db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"); if err != nil {
    log.Println(err)
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