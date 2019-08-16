package controllers

import (
  "net/http"
  "strconv"

  "alta-last/models"
  "github.com/labstack/echo"
  // "crypto/md5"
  // "fmt"
)

// get all users
func GetUsersController(c echo.Context) error {
  users, err := models.GetUsers()
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, users)
}


// get single user by id
func GetUserController(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := models.GetUser(id)

  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, user)
}


// create new user
func CreateUserController(c echo.Context) error {
  user := models.User{}
  c.Bind(&user)
  // user.Password = ToMD5(user.Password)

  result, err := models.CreateUser(&user)
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusCreated, result)
}


// remove user by id
func DeleteUserController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
  	_, err := models.DeleteUser(id)

	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success delete user",
	  
	})

}

func UpdateUserController(c echo.Context) error {
	newUser := models.User{}
  	c.Bind(&newUser)
  
	id, _ := strconv.Atoi(c.Param("id"))
  	_, err := models.UpdateUser(id, newUser)

	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success update user",
	})
}


// get all users
func GetweatherController(c echo.Context) error {
  users, err := models.GetUsers()
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, users)
}


// func ToMD5(s string) string {
//   return fmt.Sprintf("%x", md5.Sum([]byte(s)))
// }