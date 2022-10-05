package database

import (
	"echo-book/model"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = "root"      //util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = "root123"   //util.GetConfig("DB_PASSWORD")
	DB_NAME     string = "echo_book" //util.GetConfig("DB_NAME")
	DB_HOST     string = "localhost" //util.GetConfig("DB_HOST")
	DB_PORT     string = "3306"      //util.GetConfig("DB_PORT")
)

func Connect() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&model.User{}, &model.Book{})
}

type Config struct {
	DB_Username  string
	DB_Password  string
	DB_Port      string
	DB_Host      string
	DB_Test_Name string
}

func InitTestDB() {
	var err error

	config := Config{
		DB_Username:  "root",
		DB_Password:  "root123",
		DB_Port:      "3306",
		DB_Host:      "localhost",
		DB_Test_Name: "echo_book_test",
	}

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Test_Name,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")
	DB.AutoMigrate(&model.User{}, &model.Book{})
}

func SeedBook() model.Book {

	var book model.Book = model.Book{
		JudulBuku:   " Test",
		Kategori:    "Testing",
		Pengarang:   " test",
		Penerbit:    " test",
		TahunTerbit: " 2020",
		Isbn:        121311414,
	}

	if err := DB.Create(&book).Error; err != nil {
		panic(err)
	}

	var createdBook model.Book

	DB.Last(&createdBook)

	return createdBook
}

func SeedUser() model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user model.User = model.User{
		Name:     "Testing",
		Email:    "testing@mail.com",
		Password: string(password),
	}

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser model.User

	DB.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}
