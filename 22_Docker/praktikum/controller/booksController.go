package controller

import (
	"net/http"
	"strconv"

	"echo-book/config"
	"echo-book/model"

	"github.com/labstack/echo"
)

// get all Books
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all book",
		"users":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	// your solution here
	var book model.Book

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&book, "id = ?", id)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "Book not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"user":    book,
	})
}

// create new user
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"user":    book,
	})
}

// delete user by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	var book model.Book
	c.Bind(&book)

	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.First(&book, "id = ?", id)
	config.DB.Delete(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update user by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	book := model.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.First(&book, id)

	if book.ID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "invalid id",
		})
	}

	config.DB.Model(&book).Where("id= ?", id).Update(book.JudulBuku, book.Kategori, book.Pengarang, book.Penerbit, book.Isbn)

	c.Bind(&book)

	config.DB.Save(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})

}
