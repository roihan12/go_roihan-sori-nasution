package main

import (
	_driverFactory "echo-blog/drivers"
	"echo-blog/util"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	_blogUseCase "echo-blog/bussiness/blog"
	_blogController "echo-blog/controllers/blog"
	_blogRequest "echo-blog/controllers/blog/request"

	"echo-blog/controllers/users/request"

	_categoryUseCase "echo-blog/bussiness/category"
	_categoryController "echo-blog/controllers/category"

	_userUseCase "echo-blog/bussiness/users"
	_userController "echo-blog/controllers/users"

	_dbDriver "echo-blog/drivers/mysql"
	"echo-blog/drivers/mysql/blog"
	"echo-blog/drivers/mysql/category"
	"echo-blog/drivers/mysql/users"

	_middleware "echo-blog/app/middlewares"
	_routes "echo-blog/app/routes"

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

func getBlog() blog.Blog {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	note := _dbDriver.SeedBlog(db)

	return note
}

func getCategory() category.Category {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	category := _dbDriver.SeedCategory(db)

	return category
}

func TestSignup_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "test@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/users/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestSignup_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/signup").
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

func TestGetBlogs_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/blog").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetBlog_Success(t *testing.T) {
	var blog blog.Blog = getBlog()

	blogID := strconv.Itoa(int(blog.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/blog/"+blogID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetBlog_NotFound(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/api/v1/blog/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateBlog_Success(t *testing.T) {
	category := getCategory()

	var blogRequest *_blogRequest.Blog = &_blogRequest.Blog{
		Title:       "test",
		Description: "test",
		Author:      "test",
		CategoryID:  category.ID,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/blog").
		Header("Authorization", token).
		JSON(blogRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateBlog_ValidationFailed(t *testing.T) {
	var blogRequest *blog.Blog = &blog.Blog{}

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/blog").
		Header("Authorization", token).
		JSON(blogRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateBlog_Success(t *testing.T) {
	var blogs blog.Blog = getBlog()

	category := getCategory()

	var blogRequest *_blogRequest.Blog = &_blogRequest.Blog{
		Title:       "testing",
		Description: "testing",
		Author:      "testing",
		CategoryID:  category.ID,
	}

	blogID := strconv.Itoa(int(blogs.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/api/v1/blog/"+blogID).
		Header("Authorization", token).
		JSON(blogRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateBlog_ValidationFailed(t *testing.T) {
	var blogs blog.Blog = getBlog()

	var blogRequest *_blogRequest.Blog = &_blogRequest.Blog{}

	blogID := strconv.Itoa(int(blogs.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Put("/api/v1/blog/"+blogID).
		Header("Authorization", token).
		JSON(blogRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteBlog_Success(t *testing.T) {
	var blog blog.Blog = getBlog()

	var token string = getJWTToken(t)

	blogID := strconv.Itoa(int(blog.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/api/v1/blog/"+blogID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteBlog_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Delete("/api/v1/blog/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
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
