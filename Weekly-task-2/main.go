package main

import (
	"echo-item/database"
	"echo-item/route"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	server := echo.New()

	route.SetupRoute(server)

	server.Logger.Fatal(server.Start(":1323"))
}
