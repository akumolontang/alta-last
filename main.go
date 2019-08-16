package main

import (
  "alta-last/config"
  "alta-last/models"
  "alta-last/routes"

   m "alta-last/middlewares"
)

func main() {
  e := routes.New()

  // implement middleware
    m.LogMiddlewares(e)
  // migration database
  
  InitialMigration()

  e.Start(":8000")
}

func InitialMigration() {
  var db = config.DB
  if !db.HasTable(&models.User{}) { // check database exist or not
    db.AutoMigrate(&models.User{})
  }
}
