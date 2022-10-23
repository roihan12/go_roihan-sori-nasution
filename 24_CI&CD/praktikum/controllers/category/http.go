package categories

import (
	"echo-blog/bussiness/category"
	controller "echo-blog/controllers"
	"echo-blog/controllers/category/request"
	"echo-blog/controllers/category/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase category.Usecase
}

func NewCategoryController(categoryUC category.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAllCategories(c echo.Context) error {
	categoriesData := ctrl.categoryUseCase.GetAll()

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (ctrl *CategoryController) CreateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category := ctrl.categoryUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (ctrl *CategoryController) UpdateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var id string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category := ctrl.categoryUseCase.Update(id, input.ToDomain())

	return controller.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (ctrl *CategoryController) DeleteCategory(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := ctrl.categoryUseCase.Delete(id)

	if !isDeleted {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
