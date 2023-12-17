package authhandler

import (
	"gofiber-restapi/database"
	"gofiber-restapi/helper"
	"gofiber-restapi/internals/model"
	"strings"

	jtoken "github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	db := database.DB.Db
	auth := new(model.Auth)
	validate := validator.New()

	// parser body
	if err := ctx.BodyParser(auth); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Reveiw your input")
	}

	// validate
	err := validate.Struct(auth)
	if exception, ok := err.(validator.ValidationErrors); ok {
		return fiber.NewError(fiber.StatusBadRequest, exception.Error())
	}

	// verify user credentials
	var user model.User
	if err := db.Where("username = ?", auth.Username).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Username or Password incorrect")
	}

	// match password
	if !helper.MatchPassword(user.Password, auth.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Username or Password incorrect")
	}

	// generate tokens
	accessToken, refreshToken, err := helper.GenerateTokens(user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// update refresh token to database
	if err := db.Model(&user).Update("token", refreshToken).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update refresh token")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login success",
		"data": map[string]interface{}{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func RefreshToken(ctx *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	// get refresh token
	getToken := ctx.Request().Header.Peek("Authorization")
	refreshToken := strings.TrimPrefix(string(getToken), "Bearer ")

	// get data from token
	claims := ctx.Locals("user").(*jtoken.Token).Claims.(jtoken.MapClaims)

	// get user
	if err := db.Where("username = ?", claims["username"]).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Access denied")
	}

	// verify refresh token
	if user.Token != refreshToken {
		return fiber.NewError(fiber.StatusUnauthorized, "Access denied")
	}

	// generate tokens
	accessToken, refreshToken, err := helper.GenerateTokens(user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// update refresh token to database
	if err := db.Model(&user).Update("token", refreshToken).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update refresh token")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Refresh token success",
		"data": map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func Logout(ctx *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	// get data from token
	claims := ctx.Locals("user").(*jtoken.Token).Claims.(jtoken.MapClaims)

	// find user
	if err := db.Model(&user).Where("username = ?", claims["username"]).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Access denied")
	}

	// delete token
	if err := db.Model(&user).Update("token", nil).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete refresh token")
	}
	

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logout success",
		"data":    nil,
	})
}
