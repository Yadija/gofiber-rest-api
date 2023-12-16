package authroutes

import (
	authhandler "gofiber-restapi/internals/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/", authhandler.Login)
}