package routes

import (
	"echo-blog/controllers/blog"
	categories "echo-blog/controllers/category"

	"echo-blog/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      middleware.JWTConfig
	AuthController     users.AuthController
	BlogController     blog.BlogController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/signup", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	blog := e.Group("/api/v1/blog", middleware.JWTWithConfig(cl.JWTMiddleware))

	blog.GET("", cl.BlogController.GetAll)
	blog.GET("/:id", cl.BlogController.GetByID)
	blog.GET("/categories/:id", cl.BlogController.GetByCategoryID)
	blog.POST("", cl.BlogController.Create)
	blog.PUT("/:id", cl.BlogController.Update)
	blog.DELETE("/:id", cl.BlogController.Delete)

	category := e.Group("/api/v1/categories", middleware.JWTWithConfig(cl.JWTMiddleware))

	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PUT("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)

}
