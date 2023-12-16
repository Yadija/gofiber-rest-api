package router

import (
	userroutes "gofiber-restapi/internals/routes/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userroutes.SetupRoutes(app)
}
