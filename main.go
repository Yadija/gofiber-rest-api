package main

import (
	"gofiber-restapi/database"
	"gofiber-restapi/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connection to database
	database.OpenConnection()

	// setup router
	router.SetupRoutes(app)

	app.Listen(":3000")
}
