package services

import (
	"echo-book/database"
	"echo-book/model"
	"errors"
)

func GetAllBooks() ([]model.Book, error) {

	var books []model.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return []model.Book{}, err
	}

	return books, nil
}

func GetBookByID(id string) (model.Book, error) {
	var book model.Book = model.Book{}

	if err := database.DB.Find(&book, "id = ?", id).Error; err != nil {
		return model.Book{}, err
	}

	if book.ID == 0 {
		return model.Book{}, errors.New("user not found")
	}

	return book, nil
}

func CreateBook(input model.Book) (model.Book, error) {
	var book model.Book = model.Book{
		JudulBuku:   input.JudulBuku,
		Kategori:    input.Kategori,
		Pengarang:   input.Pengarang,
		Penerbit:    input.Penerbit,
		TahunTerbit: input.TahunTerbit,
		Isbn:        input.Isbn,
	}

	if err := database.DB.Save(&book).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil

}

func UpdateBook(input model.Book, id string) (model.Book, error) {
	book, err := GetBookByID(id)

	if err != nil {
		return model.Book{}, nil
	}

	book.JudulBuku = input.JudulBuku
	book.Kategori = input.Kategori
	book.Pengarang = input.Pengarang
	book.Penerbit = input.Penerbit
	book.TahunTerbit = input.TahunTerbit
	book.Isbn = input.Isbn

	if err := database.DB.Save(&book).Error; err != nil {
		return model.Book{}, nil
	}

	return book, nil
}

func DeleteBook(id string) bool {
	book, err := GetBookByID(id)

	if err != nil {
		return false
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return false
	}

	return true
}
