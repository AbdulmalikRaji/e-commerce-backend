package middleware

import (
	"time"

	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/abdulmalikraji/e-commerce/utils/messages"
	"github.com/gofiber/fiber/v2"
)

const (
	bearerPrefix = "Bearer "
)

// Helper function to refresh token and update headers
func refreshTokenAndUpdateHeaders(c *fiber.Ctx, authService services.AuthService) error {
	refreshToken := authService.GetRefreshToken(c)
	newToken, err := authService.RefreshToken(c, refreshToken)
	if err != nil {
		return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized,
			messages.CreateMsg(c, messages.InvalidToken, nil))
	}

	// Update headers with new token
	c.Set("Authorization", bearerPrefix+newToken)
	c.Request().Header.Set("Authorization", bearerPrefix+newToken)
	return nil
}

// Helper function to clean bearer prefix from token
func cleanBearerPrefix(token string) string {
	if len(token) > len(bearerPrefix) && token[:len(bearerPrefix)] == bearerPrefix {
		return token[len(bearerPrefix):]
	}
	return token
}

// TokenValidationMiddleware checks if the token is valid and refreshes if needed
func TokenValidationMiddleware(authService services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized,
				messages.CreateMsg(c, messages.RequiredField, map[string]string{"Field": "authorization token"}))
		}

		// Clean the token
		token := cleanBearerPrefix(authHeader)

		// First check if we already have a valid token in context
		if existingToken := authService.GetAccessToken(c); existingToken != "" {
			// Only try to refresh if it's expired
			if authService.IsTokenExpired(c) {
				if err := refreshTokenAndUpdateHeaders(c, authService); err != nil {
					return err
				}
			}
			return c.Next()
		}

		// Validate the new token
		valid, err := authService.ValidateToken(c, token)
		if err != nil || !valid {
			// Token is invalid, try to refresh if we have a refresh token
			if refreshToken := authService.GetRefreshToken(c); refreshToken != "" {
				if err := refreshTokenAndUpdateHeaders(c, authService); err != nil {
					return err
				}
				return c.Next()
			}
			return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized,
				messages.CreateMsg(c, messages.InvalidToken, nil))
		}

		// Store the valid token in context
		authService.SetTokens(c, token, token, time.Now().Add(1*time.Hour).Unix())
		return c.Next()
	}
}
