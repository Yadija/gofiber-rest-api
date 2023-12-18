package main

import (
	"gofiber-restapi/database"
	"gofiber-restapi/router"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError // status code defaults to 500
			// retrieve the status code if it's an *fiber.Error
			if exception, ok := err.(*fiber.Error); ok {
				code = exception.Code
			}

			// send custom errors as JSON
			return ctx.Status(code).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
				"data":    nil,
			})
		},
	})

	// connection to database
	database.OpenConnection()

	// swagger
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./apidoc.json",
		Path:     "/swagger",
		Title:    "Gofiber swagger",
	}))

	// setup router
	router.SetupRoutes(app)

	app.Listen(":3000")
}
