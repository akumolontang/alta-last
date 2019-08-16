package routes

import (
  c "alta-last/controllers"
  auth "alta-last/middlewares"
  m "github.com/labstack/echo/middleware"
  "github.com/labstack/echo"
  "alta-last/models"
  
)

func New() *echo.Echo {
  e := echo.New()

  // user routing
  e.GET("/api/users", c.GetUsersController, m.BasicAuth(auth.BasicAuth))
  e.GET("/api/users/:id", c.GetUserController, m.BasicAuth(auth.BasicAuth))
  // e.POST("/api/users", c.CreateUserController)
  e.DELETE("/api/users/:id", c.DeleteUserController, m.BasicAuth(auth.BasicAuth))
  e.PUT("/api/users/:id", c.UpdateUserController, m.BasicAuth(auth.BasicAuth))
  e.GET("api/weather/:city", models.GetWeatherController, m.BasicAuth(auth.BasicAuth))
  return e
}



