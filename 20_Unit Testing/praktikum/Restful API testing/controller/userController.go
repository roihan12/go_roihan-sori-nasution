package controller

import (
	"net/http"

	"echo-book/model"

	"echo-book/services"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := services.GetAllusers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var userId string = c.Param("id")

	user, err := services.GetUserByID(userId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"messages": "User not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get user",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	userInput := model.User{}
	c.Bind(&userInput)

	user, err := services.CreateUser(userInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here

	var userId string = c.Param("id")

	isDeleted := services.DeleteUser(userId)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "user not found",
		})

	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	var userId string = c.Param("id")

	var userInput model.User = model.User{}
	c.Bind(&userInput)

	user, err := services.UpdateUser(userInput, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})

}
