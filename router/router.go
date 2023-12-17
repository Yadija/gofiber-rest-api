package router

import (
	authroutes "gofiber-restapi/internals/routes/auth"
	threadroutes "gofiber-restapi/internals/routes/thread"
	userroutes "gofiber-restapi/internals/routes/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userroutes.SetupRoutes(app)
	authroutes.SetupRoutes(app)
	threadroutes.SetupRoutes(app)
}
