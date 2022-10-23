package route

import (
	"echo-item/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoute(server *echo.Echo) {
	// routes for item
	server.GET("/api/v1/item", controller.GetAll)
	server.GET("/api/v1/item/:id", controller.GetByID)
	server.GET("/api/v1/item/category/:id", controller.GetByCategoryID)
	server.POST("/api/v1/item", controller.Create)
	server.PUT("/api/v1/item/:id", controller.Update)
	server.DELETE("/api/v1/item/:id", controller.Delete)

	// routes for categories
	server.GET("/api/v1/categories", controller.GetAllCategories)
	server.POST("/api/v1/categories", controller.CreateCategory)
	server.PUT("/api/v1/categories/:id", controller.UpdateCategory)
	server.DELETE("/api/v1/categories/:id", controller.DeleteCategory)
}
