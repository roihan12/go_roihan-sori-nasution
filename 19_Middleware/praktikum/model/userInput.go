package model

import "github.com/go-playground/validator/v10"

type UserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (input *UserInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

func (input *UserRegister) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
