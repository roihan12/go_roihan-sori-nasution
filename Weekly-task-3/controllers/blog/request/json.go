package request

import (
	"echo-blog/bussiness/blog"

	"github.com/go-playground/validator/v10"
)

type Blog struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

func (req *Blog) ToDomain() *blog.Domain {
	return &blog.Domain{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		CategoryID:  req.CategoryID,
	}
}

func (req *Blog) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
