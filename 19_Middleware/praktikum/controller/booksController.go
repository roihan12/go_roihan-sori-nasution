package controller

import (
	"echo-book/model"
	"echo-book/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all Books
func GetBooksController(c echo.Context) error {
	books, err := services.GetAllBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all book",
		"Books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	// your solution here

	var bookId string = c.Param("id")

	book, err := services.GetBookByID(bookId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"messages": "Book not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get Book",
		"user":    book,
	})
}

// create new Book
func CreateBookController(c echo.Context) error {
	bookInput := model.Book{}
	c.Bind(&bookInput)

	book, err := services.CreateBook(bookInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new book",
		"user":    book,
	})
}

// delete Book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	var bookId string = c.Param("id")

	isDeleted := services.DeleteUser(bookId)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "book not found",
		})

	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete book",
	})
}

// update Book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	var bookId string = c.Param("id")

	var bookInput model.Book = model.Book{}
	c.Bind(&bookInput)

	book, err := services.UpdateBook(bookInput, bookId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update book",
		"user":    book,
	})

}
