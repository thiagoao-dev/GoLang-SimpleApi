package main

import (
  "time"
  
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Name       string     `sql:"type:varchar(50);index;not null" json:"name"`
  Birthday   time.Time  `sql:"type:date;not null" json:"birthday"`
  Username   string     `sql:"type:varchar(20):not null;unique_index" json:"username"`
  Phones     []Phone    `json:"phones"`
  Emails     []Email    `json:"emails"`
  Languages  []Language `gorm:"many2many:user_languages;" json:"languages"`
}

type Phone struct {
  gorm.Model
  Prefix     string `sql:"type:varchar(3);not null" json:"phone_prefix"`
  Number     string `sql:"type:varchar(10);not null" json:"phone_number"`
}

type Email struct {
  gorm.Model
  UserID     int    `sql:"index"`
  Email      string `sql:"type:varchar(100);not null" json:"email"`
  Primary    bool   `sql:"type:boolean" json:"primary"`
}

type Language struct {
  ID   int
  Name string `sql:"unique_index:idx_name_code" json:"lang_name"`
  Code string `sql:"index:idx_name_code" json:"lang_code"`
}