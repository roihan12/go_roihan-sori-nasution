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

func InitEcho() *echo.Echo {
	database.InitTestDB()
	e := echo.New()

	return e
}

func TestGetAllBooks_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
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

func TestCreateBook_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/books",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	bookInput := model.Book{
		JudulBuku:   " Test",
		Kategori:    "Testing",
		Pengarang:   " test",
		Penerbit:    " test",
		TahunTerbit: " 2020",
		Isbn:        121311414,
	}

	jsonBody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/books", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestGetBookByID_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	book := database.SeedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestUpdateBook_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	book := database.SeedBook()

	bookInput := model.Book{
		JudulBuku:   " Test",
		Kategori:    "Testing",
		Pengarang:   " test",
		Penerbit:    " test",
		TahunTerbit: " 2020",
		Isbn:        121311414,
	}

	jsonBody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonBody)

	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodPut, "/books", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestDeleteBook_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	},
	}

	e := InitEcho()

	book := database.SeedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}
