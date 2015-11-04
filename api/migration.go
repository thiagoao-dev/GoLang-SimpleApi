package api

import (
  "github.com/jinzhu/gorm"
)

func migrationUp() error {
  
  // Drop table
  db.DropTable(&User{}, &Phone{}, &Email{}, &Language{})
  
  // Create table
  //db.CreateTable(&User{})
  db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}, &Phone{}, &Email{}, &Language{})

  // ModifyColumn
  //db.Model(&User{}).ModifyColumn("description", "text")

  // DropColumn
  //db.Model(&User{}).DropColumn("description")

  // Automating Migration
  db.AutoMigrate(&User{}, &Phone{}, &Email{}, &Language{})
}

