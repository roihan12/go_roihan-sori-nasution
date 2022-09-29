package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"echo-book/route"

	"echo-book/config"
)

func main() {
	config.InitDB()
	e := route.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":1323"))
}
