package database

import (
	"echo-book/model"
	"echo-book/util"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = util.GetConfig("DB_PASSWORD")
	DB_NAME     string = util.GetConfig("DB_NAME")
	DB_HOST     string = util.GetConfig("DB_HOST")
	DB_PORT     string = util.GetConfig("DB_PORT")
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
