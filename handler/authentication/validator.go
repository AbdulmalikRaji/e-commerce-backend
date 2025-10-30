package authentication

import (
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/abdulmalikraji/e-commerce/utils"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/abdulmalikraji/e-commerce/utils/messages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func LoginByEmailRequestValidator(ctx *fiber.Ctx) error {

	var request authDto.LoginByEmailRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Error(err)
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	// Username field could be email or phone number
	// validate using with regex
	if request.Email == "" {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Email"}))
	}

	if request.Password == "" {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, messages.CreateMsg(ctx, messages.RequiredField, map[string]string{"Field": "Password"}))
	}

	if !utils.EmailRegex(request.Email) {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, messages.CreateMsg(ctx, messages.InvalidFormat, map[string]string{"Field": "Email"}))
	}

	return ctx.Next()
}
