package main

import (
	"gofiber-restapi/database"
	"gofiber-restapi/router"

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
				"status":  code,
				"message": err.Error(),
			})
		},
	})

	// connection to database
	database.OpenConnection()

	// setup router
	router.SetupRoutes(app)

	app.Listen(":3000")
}
