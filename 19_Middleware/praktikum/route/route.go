package route

import (
	"echo-book/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {

	// route for auth

	server.POST("/users/register", controller.Register)
	server.POST("/users/login", controller.Login)

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))

	// Route / to handler function
	//Route Users
	server.POST("/users", controller.CreateUserController)
	server.GET("/books", controller.GetBooksController)
	server.GET("/books/:id", controller.GetBookController)

	privateRoutes.GET("/users", controller.GetUsersController)
	privateRoutes.GET("/users/:id", controller.GetUserController)
	privateRoutes.DELETE("/users/:id", controller.DeleteUserController)
	privateRoutes.PUT("/users/:id", controller.UpdateUserController)

	//Route Books
	privateRoutes.POST("/books", controller.CreateBookController)
	privateRoutes.DELETE("/books/:id", controller.DeleteBookController)
	privateRoutes.PUT("/books/:id", controller.UpdateBookController)

}
