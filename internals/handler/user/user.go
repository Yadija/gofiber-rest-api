package userhandler

import (
	"gofiber-restapi/database"
	"gofiber-restapi/helper"
	"gofiber-restapi/internals/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
	validate := validator.New()

	// parse body
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Reveiw your input",
		})
	}

	err := validate.Struct(user)
	if exception, ok := err.(validator.ValidationErrors); ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": exception.Error(),
		})
	}

	// check if username already exist
	if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Username already exist",
		})
	}

	// hashing password
	hashedPassword, _ := helper.HashedPassword(user.Password)
	user.Password = hashedPassword

	// create user
	err = db.Create(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Couldn't create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data": map[string]interface{}{
			"username": user.Username,
			"fullname": user.Fullname,
		},
	})
}
