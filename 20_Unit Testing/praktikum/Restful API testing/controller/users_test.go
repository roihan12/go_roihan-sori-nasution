package controller

import (
	"bytes"
	"echo-book/database"
	"echo-book/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InittEcho() *echo.Echo {
	database.InitTestDB()
	e := echo.New()

	return e
}

func TestGetAllUsers_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// body := rec.Body.String()

			// assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestLoginUsers_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"token\":",
	},
	}

	e := InitEcho()
	user := database.SeedUser()

	userInput := model.UserInput{
		Email:    user.Email,
		Password: user.Password,
	}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestLoginUsers_Failed(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"token\":",
	},
	}

	e := InitEcho()
	// user := database.SeedUser()

	userInput := model.UserInput{
		Email:    "wrong@gmail.com",
		Password: "123",
	}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		err := Login(c)

		if assert.Error(t, err) {
			expected := echo.NewHTTPError(http.StatusUnauthorized)
			assert.Equal(t, http.StatusOK, expected, err)
		}
	}

}

func TestCreateUser_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	user := database.SeedUser()

	userRegister := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	jsonBody, _ := json.Marshal(&userRegister)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestGetUserByID_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	user := database.SeedUser()
	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userID)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestUpdateUser_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	user := database.SeedUser()

	userInput := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodPut, "/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userID)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestDeleteUser_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	user := database.SeedUser()
	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userID)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}
