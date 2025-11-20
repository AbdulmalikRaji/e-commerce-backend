package middleware

import (
	"strings"

	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/abdulmalikraji/e-commerce/utils/messages"
	"github.com/gofiber/fiber/v2"
)

const bearerPrefix = "Bearer "

// extractToken extracts the token from the Authorization header
func extractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing authorization header")
	}

	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization header format")
	}

	return authHeader[len(bearerPrefix):], nil
}

// TokenValidationMiddleware checks if the token is valid and refreshes if needed
func TokenValidationMiddleware(authService services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip auth for login and refresh routes
		if c.Path() == "/auth/login" || c.Path() == "/auth/refresh" {
			return c.Next()
		}

		// Extract token from header
		token, err := extractToken(c.Get("Authorization"))
		if err != nil {
			return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized,
				messages.CreateMsg(c, messages.RequiredField, map[string]string{"Field": "authorization token"}))
		}

		// Validate token
		if err := authService.ValidateToken(c, token); err != nil {
			// Token is invalid, try to refresh
			refreshData, err := authService.RefreshToken(c)
			if err != nil {
				return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized,
					messages.CreateMsg(c, messages.InvalidToken, nil))
			}
			
			// Update request header for downstream handlers
			newToken := refreshData.AccessToken
			c.Request().Header.Set("Authorization", bearerPrefix+newToken)
		}

		return c.Next()
	}
}
