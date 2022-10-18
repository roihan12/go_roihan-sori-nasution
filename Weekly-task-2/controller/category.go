package controller

import (
	"echo-item/model"
	"echo-item/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var categoryService service.CategoryService = service.NewCategoryService()

func GetAllCategories(c echo.Context) error {
	var categories []model.Category = categoryService.GetAll()

	return c.JSON(http.StatusOK, model.Response[[]model.Category]{
		Status:  "success",
		Message: "all categories",
		Data:    categories,
	})
}

func CreateCategory(c echo.Context) error {
	var input *model.CategoryInput = new(model.CategoryInput)

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    nil,
		})
	}

	category := categoryService.Create(*input)

	return c.JSON(http.StatusCreated, model.Response[model.Category]{
		Status:  "success",
		Message: "category created",
		Data:    category,
	})
}

func UpdateCategory(c echo.Context) error {
	var input *model.CategoryInput = new(model.CategoryInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    nil,
		})
	}

	var id string = c.Param("id")

	category := categoryService.Update(id, *input)

	return c.JSON(http.StatusOK, model.Response[model.Category]{
		Status:  "success",
		Message: "category updated",
		Data:    category,
	})
}

func DeleteCategory(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := categoryService.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, model.Response[bool]{
			Status:  "failed",
			Message: "category not found",
			Data:    isDeleted,
		})
	}

	return c.JSON(http.StatusOK, model.Response[bool]{
		Status:  "success",
		Message: "category deleted",
		Data:    isDeleted,
	})
}
