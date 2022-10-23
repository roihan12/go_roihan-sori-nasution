package main

import (
	_driverFactory "echo-blog/drivers"
	"echo-blog/util"
	"log"

	_blogUseCase "echo-blog/bussiness/blog"
	_blogController "echo-blog/controllers/blog"

	_categoryUseCase "echo-blog/bussiness/category"
	_categoryController "echo-blog/controllers/category"

	_userUseCase "echo-blog/bussiness/users"
	_userController "echo-blog/controllers/users"

	_dbDriver "echo-blog/drivers/mysql"

	_middleware "echo-blog/app/middlewares"
	_routes "echo-blog/app/routes"

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

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUseCase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	blogRepo := _driverFactory.NewBlogRepository(db)
	blogUsecase := _blogUseCase.NewBlogUsecase(blogRepo)
	blogCtrl := _blogController.NewBlogController(blogUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		BlogController:     *blogCtrl,
		AuthController:     *userCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":1323"))
}
