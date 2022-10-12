package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Book struct {
	id       string
	judul    string
	price    int
	kategori string
}

var books []Book

func listMenu() {
	fmt.Println("===BOOK MANAGEMENT===")
	fmt.Println("1. Get all books")
	fmt.Println("2. Get books by id")
	fmt.Println("3. Create a book")
	fmt.Println("4. Update a book")
	fmt.Println("5. Delete a book")
	fmt.Println("6. Exit")
	fmt.Print("Choose your menu: ")
}

func menu() {
	var choose int
	listMenu()
	for {
		fmt.Scan(&choose)
		if choose == 6 {
			fmt.Println("Bye...")
			break
		} else {
			switch choose {
			case 1:
				fmt.Println("All books")
				getAllBook(books)
				fmt.Println()
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 2:
				var id string
				fmt.Println()
				fmt.Print("Enter Product ID :")
				fmt.Scan(&id)
				fmt.Println("\nInfo for Product with ID ", id)
				getBookById(id, books)
				fmt.Println()
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 3:
				fmt.Println("Add a new product")
				addBook(&books)
				fmt.Println("Book Added!")
				fmt.Println()
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 4:
				var id string
				fmt.Println()
				fmt.Print("\nEnter Product Id: ")
				fmt.Scan(&id)
				updateBookById(id, &books)
				fmt.Println("Book updated!")
				fmt.Println()
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			case 5:
				var id string
				fmt.Println()
				fmt.Print("\nEnter Product Id: ")
				fmt.Scan(&id)
				deleteBookById(id, &books)
				fmt.Println("Book deleted!")
				fmt.Println("press 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				listMenu()
			default:
				fmt.Println("Re-enter your choice!")
				listMenu()
			}
		}
	}
}

func getAllBook(books []Book) {
	for _, book := range books {
		fmt.Println("===")
		fmt.Println("ID	:", book.id)
		fmt.Println("Judul	:", book.judul)
		fmt.Println("Price	:", book.price)
		fmt.Println("Kategori:", book.kategori)
		fmt.Println("===")
	}
}

func getBookById(id string, book []Book) {
	for _, book := range book {
		if book.id == id {
			fmt.Println("===")
			fmt.Println("ID	:", book.id)
			fmt.Println("Name	:", book.judul)
			fmt.Println("Price	:", book.price)
			fmt.Println("Kategori:", book.kategori)
			fmt.Println("===")
		}
	}
}

func addBook(books *[]Book) {

	uuid := uuid.NewString()
	id := uuid

	reader := bufio.NewReader(os.Stdin)
	var price int
	var kategori string
	fmt.Println()
	fmt.Print("Enter Title : ")
	judul, _ := reader.ReadString('\n')

	fmt.Scan(&judul)
	fmt.Print("Enter Price : ")
	fmt.Scan(&price)

	fmt.Print("Enter Kategori : ")
	fmt.Scan(&kategori)

	book := Book{id, judul, price, kategori}

	*books = append(*books, book)
}

func updateBookById(id string, books *[]Book) {

	var judul string
	var price int
	var kategori string

	fmt.Println()
	fmt.Print("Enter the New title : ")
	fmt.Scan(&judul)

	fmt.Print("Enter the New price : ")
	fmt.Scan(&price)

	fmt.Print("Enter the New kategori : ")
	fmt.Scan(&kategori)

	for i, book := range *books {
		if book.id == id {
			(*books)[i] = Book{id, judul, price, kategori}
		}
	}
}

func deleteBookById(id string, books *[]Book) {
	for i, book := range *books {
		if book.id == id {
			*books = append((*books)[:i], (*books)[i+1:]...)
		}
	}
}

func BookApp() {

	// Example data
	books = append(books, Book{id: uuid.NewString(), judul: "Algoritma 1", price: 25000, kategori: "tech"})
	books = append(books, Book{id: uuid.NewString(), judul: "Golang", price: 15000, kategori: "tech"})
	menu()
}

func main() {
	BookApp()
}
