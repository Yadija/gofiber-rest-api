package userroutes

import (
	userhandler "gofiber-restapi/internals/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Post("/", userhandler.CreateUser)
}