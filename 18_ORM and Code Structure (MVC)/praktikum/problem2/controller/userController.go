package controller

import (
	"net/http"
	"strconv"

	"echo-book/config"
	"echo-book/model"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
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
	var user model.User

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&user, "id = ?", id)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "User not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
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
	var user model.User
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.First(&user, "id = ?", id)
	config.DB.Delete(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	user := model.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.First(&user, id)

	if user.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}
	config.DB.Model(&user).Where("id= ?", id).Update(user.Name, user.Email, user.Password)
	c.Bind(&user)

	config.DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})

}
