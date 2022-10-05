package main

import (
	"echo-book/middlewares"
	"echo-book/route"

	"echo-book/database"

	"github.com/labstack/echo/v4"
)

func main() {
	// database.Connect()
	server := echo.New()

	database.InitTestDB()

	middlewares.LogMiddleware(server)

	route.SetupRoute(server)
	// start the server, and log if it fails
	server.Logger.Fatal(server.Start(":1323"))
}
