package routes

import (
	"praktikum/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    middleware.JWTConfig
	AuthController   users.AuthController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.CreateUser)
	users.POST("/login", cl.AuthController.Login)

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	user := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))

	user.GET("", cl.AuthController.GetAllusers)

	auth.POST("/logout", cl.AuthController.Logout)

}
