package blog

import (
	"echo-blog/bussiness/blog"
	controller "echo-blog/controllers"
	"echo-blog/controllers/blog/request"
	"echo-blog/controllers/blog/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	blogUseCase blog.Usecase
}

func NewBlogController(blogUC blog.Usecase) *BlogController {
	return &BlogController{
		blogUseCase: blogUC,
	}
}

func (ctrl *BlogController) GetByCategoryID(c echo.Context) error {
	var id string = c.Param("id")

	blogCategory := ctrl.blogUseCase.GetByCategoryID(id)

	if id == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "post not found", "")
	}

	blogs := []response.Blog{}

	for _, blog := range blogCategory {
		blogs = append(blogs, response.FromDomain(blog))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all blog", blogs)
}

func (ctrl *BlogController) GetAll(c echo.Context) error {

	keyword := c.QueryParam("keyword")

	blogData := ctrl.blogUseCase.GetAll(keyword)

	blogs := []response.Blog{}

	for _, blog := range blogData {
		blogs = append(blogs, response.FromDomain(blog))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all blog", blogs)
}

func (ctrl *BlogController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	blog := ctrl.blogUseCase.GetByID(id)

	if blog.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "post not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "post found", response.FromDomain(blog))
}

func (ctrl *BlogController) Create(c echo.Context) error {
	input := request.Blog{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	blog := ctrl.blogUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "Blog created", response.FromDomain(blog))
}

func (ctrl *BlogController) Update(c echo.Context) error {
	input := request.Blog{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var noteId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	blog := ctrl.blogUseCase.Update(noteId, input.ToDomain())

	if blog.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "Blog updated", response.FromDomain(blog))
}

func (ctrl *BlogController) Delete(c echo.Context) error {
	var blogId string = c.Param("id")

	isSuccess := ctrl.blogUseCase.Delete(blogId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "Blog not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "Blog deleted", "")
}
