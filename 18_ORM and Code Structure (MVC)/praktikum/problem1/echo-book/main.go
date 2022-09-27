package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)
 
var (
  DB *gorm.DB
)
 
func init() {
  InitDB()
  InitialMigration()
}
 
type Config struct {
  DB_Username string
  DB_Password string
  DB_Port     string
  DB_Host     string
  DB_Name     string
}
 
func InitDB() {
 
  config := Config{
    DB_Username: "root",
    DB_Password: "root123",
    DB_Port:     "3306",
    DB_Host:     "localhost",
    DB_Name:     "crud_go",
  }
 
  connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    config.DB_Username,
    config.DB_Password,
    config.DB_Host,
    config.DB_Port,
    config.DB_Name,
  )
 
  var err error
  DB, err = gorm.Open("mysql", connectionString)
  if err != nil {
    panic(err)
  }
}
 
type User struct {
  gorm.Model
  Name     string `json:"name" form:"name"`
  Email    string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}
 
func InitialMigration() {
  DB.AutoMigrate(&User{})
}
 
// get all users
func GetUsersController(c echo.Context) error {
  var users []User
 
  if err := DB.Find(&users).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success get all users",
    "users":   users,
  })
}
 
// get user by id
func GetUserController(c echo.Context) error {
  // your solution here
    var user User

    id, _ := strconv.Atoi(c.Param("id"))
	
	DB.First(&user, "id = ?", id)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "User not Found",
		})
	}


	return c.JSON(http.StatusOK, user)
}
 
// create new user
func CreateUserController(c echo.Context) error {
  user := User{}
  c.Bind(&user)
 
  if err := DB.Save(&user).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success create new user",
    "user":    user,
  })
}
 

// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here
  	var user User
  	c.Bind(&user)
  
	id, _ := strconv.Atoi(c.Param("id"))
	DB.First(&user, "id = ?", id)
	DB.Delete(&user)
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
 
// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here

  	user := User{}
 
	id, _ := strconv.Atoi(c.Param("id"))
	DB.First(&user, id)

	if user.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}
	DB.Model(&user).Where("id= ?", id).Update(user.Name,user.Email,user.Password)
	c.Bind(&user)

	DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})

}
 
func main() {
  // create a new echo instance
  e := echo.New()
  // Route / to handler function
  e.GET("/users", GetUsersController)
  e.GET("/users/:id", GetUserController)
  e.POST("/users", CreateUserController)
  e.DELETE("/users/:id", DeleteUserController)
  e.PUT("/users/:id", UpdateUserController)
 
  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}

