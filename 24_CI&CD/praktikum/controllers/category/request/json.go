package request

import (
	"echo-blog/bussiness/category"

	"github.com/go-playground/validator/v10"
)

type Category struct {
	Name string `json:"name" validate:"required"`
}

func (req *Category) ToDomain() *category.Domain {
	return &category.Domain{
		Name: req.Name,
	}
}

func (req *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
