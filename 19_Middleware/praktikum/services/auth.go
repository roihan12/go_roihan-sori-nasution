package services

import (
	"echo-book/model"
	"echo-book/repository"
)

type AuthService struct {
	Repository repository.AuthRepository
}

func NewAuthService() AuthService {
	return AuthService{
		Repository: &repository.AuthRepositoryImpl{},
	}
}

func (a *AuthService) Register(input model.UserRegister) model.User {
	return a.Repository.Register(input)
}

func (a *AuthService) Login(input model.UserInput) string {
	return a.Repository.Login(input)
}
