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
	ValidateToken(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
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

	data, status, err := c.service.LoginByEmail(ctx, loginRequest)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, status, data, "Login successful")
}

func (c authHandler) SignUp(ctx *fiber.Ctx) error {
	return nil
}

// ValidateToken validates the current token
func (c authHandler) ValidateToken(ctx *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := ctx.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	token := authHeader[7:]
	if err := c.service.ValidateToken(ctx, token); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid token")
	}

	return genericResponse.SuccessResponse(ctx, fiber.StatusOK, nil, "Token is valid")
}

// RefreshToken refreshes the current access token using the refresh token
func (c authHandler) RefreshToken(ctx *fiber.Ctx) error {
	data, err := c.service.RefreshToken(ctx)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Failed to refresh token")
	}

	return genericResponse.SuccessResponse(ctx, fiber.StatusOK, data, "Token refreshed successfully")
}

func (c authHandler) Logout(ctx *fiber.Ctx) error {
	var logoutRequest authDto.LogoutRequest

	authHeader := ctx.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	token := authHeader[7:]
	if err := c.service.ValidateToken(ctx, token); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid token")
	}

	logoutRequest.AccessToken = token

	if err := c.service.Logout(ctx, logoutRequest); err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, fiber.StatusOK, nil, "Logout successful")
}
