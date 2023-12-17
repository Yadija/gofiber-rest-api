package threadhandler

import (
	"gofiber-restapi/database"
	"gofiber-restapi/internals/model"

	jtoken "github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateThread(ctx *fiber.Ctx) error {
	db := database.DB.Db
	thread := new(model.Thread)
	validate := validator.New()

	// parser body
	if err := ctx.BodyParser(thread); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Reveiw your input")
	}

	// validate
	err := validate.Struct(thread)
	if exception, ok := err.(validator.ValidationErrors); ok {
		return fiber.NewError(fiber.StatusBadRequest, exception.Error())
	}

	// get data from token
	claims := ctx.Locals("user").(*jtoken.Token).Claims.(jtoken.MapClaims)

	// create thread
	addedThread := model.Thread{
		Content: thread.Content,
		Owner:   claims["username"].(string),
	}
	if err := db.Create(&addedThread).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't create thread")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Create thread success",
		"data": map[string]interface{}{
			"threadId": addedThread.ID,
		},
	})
}

func GetAllThreads(ctx *fiber.Ctx) error {
	db := database.DB.Db
	var threads []model.Thread

	db.Select("id", "content", "owner", "created_at", "updated_at").Find(&threads)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get all threads success",
		"data": map[string]interface{}{
			"threads": threads,
		},
	})
}

func GetThreadById(ctx *fiber.Ctx) error {
	db := database.DB.Db
	var thread model.Thread
	threadId := ctx.Params("threadId")

	if err := db.Select("id", "content", "owner", "created_at", "updated_at").Where("id = ?", threadId).First(&thread).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Thread not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get thread success",
		"data": map[string]interface{}{
			"thread": thread,
		},
	})
}
