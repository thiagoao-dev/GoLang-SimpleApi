package main

import (
  "github.com/jinzhu/gorm"
)

func migrationUp() error {
  
  // Drop table
  db.DropTable(&User{}, &Phone{}, &Email{}, &Language{})
  
  // Create table
  //db.CreateTable(&User{})
  db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}, &Phone{}, &Email{}, &Language{})

  // Drop table
  db.DropTable(&User{})

  // ModifyColumn
  //db.Model(&User{}).ModifyColumn("description", "text")

  // DropColumn
  //db.Model(&User{}).DropColumn("description")

  // Automating Migration
  //db.AutoMigrate(&User{})
  //db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
  db.AutoMigrate(&User{}, &Phone{}, &Email{}, &Language{})
  // Feel free to change your struct, AutoMigrate will keep your database up-to-date.
  // AutoMigrate will ONLY add *new columns* and *new indexes*,
  // WON'T update current column's type or delete unused columns, to protect your data.
  // If the table is not existing, AutoMigrate will create the table automatically.

}

