package repository

import "echo-book/model"

type AuthRepository interface {
	Register(input model.UserRegister) model.User
	Login(input model.UserInput) string
}
