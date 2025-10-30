package authentication

import (
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignUp(ctx *fiber.Ctx) error
	LoginByEmail(ctx *fiber.Ctx) error
}

type authHandler struct {
	service services.AuthService
}

func New(service services.AuthService) AuthHandler {
	return authHandler{
		service: service,
	}
}

func (c authHandler) LoginByEmail(ctx *fiber.Ctx) error {
	var loginRequest authDto.LoginByEmailRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	token, err := c.service.LoginByEmail(ctx, loginRequest)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, token, fiber.StatusOK)
}

func (c authHandler) SignUp(ctx *fiber.Ctx) error {
	return nil
}
