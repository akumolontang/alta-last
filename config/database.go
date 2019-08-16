package config

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

var (
  DB *gorm.DB
)

func init() {
  InitDB()
}

func InitDB() {


  db, err := gorm.Open("mysql", "root:L4f4y3tt3@/go_db2?charset=utf8&parseTime=True&loc=Local")

  if err != nil {
    panic(err)
  }
  DB = db
}

