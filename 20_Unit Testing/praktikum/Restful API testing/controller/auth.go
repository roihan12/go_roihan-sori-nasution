package controller

import (
	"echo-book/auth"
	"echo-book/model"
	"echo-book/services"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var authService services.AuthService = services.NewAuthService()

func Register(c echo.Context) error {
	var userInput *model.UserRegister = new(model.UserRegister)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	user := authService.Register(*userInput)

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var userInput *model.UserInput = new(model.UserInput)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := authService.Login(*userInput)

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	isListed := auth.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "invalid token",
		})

	}

	auth.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
