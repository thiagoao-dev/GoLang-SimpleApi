package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

func getConn() *gorm.DB {
  db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
  // db, err := gorm.Open("foundation", "dbname=gorm") // FoundationDB.
  // db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
  // db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

  // Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
  db.DB()

  // Then you could invoke `*sql.DB`'s functions with it
  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)

  // Disable table name's pluralization
  // db.SingularTable(true)
  
  return db
}