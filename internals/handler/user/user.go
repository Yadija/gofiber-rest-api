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
		return fiber.NewError(fiber.StatusInternalServerError, "Reveiw your input")
	}

	err := validate.Struct(user)
	if exception, ok := err.(validator.ValidationErrors); ok {
		return fiber.NewError(fiber.StatusInternalServerError, exception.Error())
	}

	// check if username already exist
	if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Username already exist")
	}

	// hashing password
	hashedPassword, _ := helper.HashedPassword(user.Password)
	user.Password = hashedPassword

	// create user
	err = db.Create(&user).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't create user")
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

func GetUserByUsername(ctx *fiber.Ctx) error {
	db := database.DB.Db
	username := ctx.Params("username")

	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    map[string]interface{}{
			"username": user.Username,
			"fullname": user.Fullname,
		},
	})
}
