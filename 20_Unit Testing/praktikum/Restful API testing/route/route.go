package route

import (
	"echo-book/controller"
	"echo-book/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {

	// route for auth

	server.POST("/users/register", controller.Register)
	server.POST("/users/login", controller.Login)
	// logout

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))

	privateRoutes.Use(middlewares.CheckTokenMiddleware)

	// Route / to handler function
	privateRoutes.POST("/users", controller.CreateUserController)
	privateRoutes.GET("/books", controller.GetBooksController)
	privateRoutes.GET("/books/:id", controller.GetBookController)

	// Authentication
	privateRoutes.GET("/users", controller.GetUsersController)
	privateRoutes.GET("/users/:id", controller.GetUserController)
	privateRoutes.DELETE("/users/:id", controller.DeleteUserController)
	privateRoutes.PUT("/users/:id", controller.UpdateUserController)

	privateRoutes.POST("/books", controller.CreateBookController)
	privateRoutes.DELETE("/books/:id", controller.DeleteBookController)
	privateRoutes.PUT("/books/:id", controller.UpdateBookController)

	privateRoutes.POST("/users/logout", controller.Logout)

}
