package threadroutes

import (
	threadhandler "gofiber-restapi/internals/handler/thread"
	"gofiber-restapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	accessJWT := middleware.NewAuthMiddleware("ACCESS_TOKEN_KEY")

	thread := app.Group("/threads")
	thread.Get("/", threadhandler.GetAllThreads)
	thread.Get("/:threadId", threadhandler.GetThreadById)

	thread.Use(accessJWT) // set auth middleware
	thread.Post("/", threadhandler.CreateThread)
	thread.Put("/:threadId", threadhandler.UpdateThread)
	thread.Delete("/:threadId", threadhandler.DeleteThread)
}
