package users_test

import (
	"praktikum/app/middlewares"
	"praktikum/businesses/users"
	_userMock "praktikum/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepository, &middlewares.ConfigJWT{})

	usersDomain = users.Domain{
		Email:    "test@test.com",
		Password: "123123",
	}

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("CreateUser | Valid", func(t *testing.T) {
		usersRepository.On("CreateUser", &usersDomain).Return(usersDomain).Once()

		result := usersService.CreateUser(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("CreateUser | InValid", func(t *testing.T) {
		usersRepository.On("CreateUser", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.CreateUser(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &usersDomain).Return(users.Domain{}).Once()

		result := usersService.Login(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Login(&users.Domain{})

		assert.Empty(t, result)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("GetAllusers | Valid", func(t *testing.T) {
		usersRepository.On("GetAllusers").Return([]users.Domain{usersDomain}).Once()

		result := usersService.GetAllusers()

		assert.Equal(t, 1, len(result))
	})

	t.Run("GetAllusers | InValid", func(t *testing.T) {
		usersRepository.On("GetAllusers").Return([]users.Domain{}).Once()

		result := usersService.GetAllusers()

		assert.Equal(t, 0, len(result))
	})
}
