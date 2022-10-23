package controller

import (
	"echo-item/model"
	"echo-item/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var itemService service.ItemService = service.New()

func GetAll(c echo.Context) error {

	keyword := c.QueryParam("keyword")

	var item []model.Item = itemService.GetAll(keyword)

	return c.JSON(http.StatusOK, model.Response[[]model.Item]{
		Status:  "success",
		Message: "all items",
		Data:    item,
	})

}

func GetByID(c echo.Context) error {
	var id string = c.Param("id")

	item := itemService.GetByID(id)

	if item.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "item not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response[model.Item]{
		Status:  "success",
		Message: "get items",
		Data:    item,
	})

}

func GetByCategoryID(c echo.Context) error {
	var category_id string = c.Param("id")

	var item []model.Item = itemService.GetByCategoryID(category_id)

	if category_id == "" || category_id == "0" {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "kategori not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response[[]model.Item]{
		Status:  "success",
		Message: "all items",
		Data:    item,
	})

}

func Create(c echo.Context) error {
	var input *model.ItemInput = new(model.ItemInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	item := itemService.Create(*input)

	return c.JSON(http.StatusOK, model.Response[model.Item]{
		Status:  "success",
		Message: " create items",
		Data:    item,
	})

}

func Update(c echo.Context) error {
	var input *model.ItemInput = new(model.ItemInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	var itemId string = c.Param("id")

	item := itemService.Update(itemId, *input)
	return c.JSON(http.StatusOK, model.Response[model.Item]{
		Status:  "success",
		Message: "items update",
		Data:    item,
	})

}

func Delete(c echo.Context) error {
	var itemId string = c.Param("id")

	isSuccess := itemService.Delete(itemId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete a data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "data deleted",
	})
}
