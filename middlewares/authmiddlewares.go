package middlewares

import (
	"github.com/labstack/echo"
	// "crypto/md5"
	// "alta-last/controllers"
  )

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
	  return true, nil
	}
	return false, nil
  }
  
// func BasicAuth2(username, password string, c.echoContext) (bool, error) {
// 	var db = config.db
// 	var user m.User
// 	if err := db.Where("email = ? AND password = ?", username, controllers.ToMD5(password))
// 	.First(&user).Error; err != nil {
// 		return false, nil
// 	}
// 	return true, nil
// }