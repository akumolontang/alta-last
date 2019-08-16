package controllers

import (
  "net/http"
//   "strconv"

//   "alta-last/models"
  "github.com/labstack/echo"
//   "crypto/md5"
//   "fmt"
)



// get single weather by city
func GetWeatherController(c echo.Context) error {
  city := c.Param("city")

  return c.JSON(http.StatusOK, map[string]interface{}{
	"city": city,
	"temperatur": 30,
  })
}

