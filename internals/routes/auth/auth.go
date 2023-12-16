package authroutes

import (
	authhandler "gofiber-restapi/internals/handler/auth"
	"gofiber-restapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	refreshJWT := middleware.NewAuthMiddleware("REFRESH_TOKEN_KEY")

	auth := router.Group("/auth")
	auth.Post("/", authhandler.Login)
	auth.Put("/", refreshJWT, authhandler.RefreshToken)
}