package main

import (
	"echo-item/database"
	"echo-item/route"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "1323"

func main() {
	database.Connect()

	server := echo.New()

	route.SetupRoute(server)

	var port string = os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}
	var appPort string = fmt.Sprintf(":%s", port)

	server.Logger.Fatal(server.Start(appPort))
}
