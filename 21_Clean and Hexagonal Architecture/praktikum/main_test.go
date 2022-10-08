package main

import (
	"encoding/json"
	"net/http"
	_driverFactory "praktikum/drivers"
	"praktikum/util"
	"testing"

	_userUseCase "praktikum/businesses/users"
	_userController "praktikum/controllers/users"
	"praktikum/controllers/users/request"

	_dbDriver "praktikum/drivers/mysql"
	"praktikum/drivers/mysql/users"

	_middleware "praktikum/app/middlewares"
	_routes "praktikum/app/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware:    configJWT.Init(),
		AuthController:   *userCtrl,
	}

	routesInit.RouteRegister(e)

	return e
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: util.GetConfig("DB_USERNAME"),
			DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
			DB_HOST:     util.GetConfig("DB_HOST"),
			DB_PORT:     util.GetConfig("DB_PORT"),
			DB_NAME:     util.GetConfig("DB_TEST_NAME"),
		}

		db := configDB.InitDB()

		_dbDriver.CleanSeeders(db)
	}
}

func getJWTToken(t *testing.T) string {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	return user
}

func TestRegister_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "test@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := getUser()

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestLogout_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogout_Failed(t *testing.T) {
	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
