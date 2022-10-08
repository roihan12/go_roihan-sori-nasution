package main

import (
	"log"
	_driverFactory "praktikum/drivers"
	"praktikum/util"

	_userUseCase "praktikum/businesses/users"
	_userController "praktikum/controllers/users"

	_dbDriver "praktikum/drivers/mysql"

	_middleware "praktikum/app/middlewares"
	_routes "praktikum/app/routes"

	echo "github.com/labstack/echo/v4"
)

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_NAME"),
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

	log.Fatal(e.Start(":1323"))
}
