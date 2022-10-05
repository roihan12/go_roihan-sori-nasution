package repository

import (
	"echo-book/auth"
	"echo-book/database"
	"echo-book/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
}

func (a *AuthRepositoryImpl) Register(input model.UserRegister) model.User {

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser model.User = model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}

	var createdUser model.User = model.User{}

	result := database.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (a *AuthRepositoryImpl) Login(input model.UserInput) string {

	var user model.User = model.User{}

	database.DB.First(&user, "email = ?", input.Email)

	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}

	token := auth.CreateToken(user.ID)

	return token
}
