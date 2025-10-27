package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Auth0Middleware checks if the user is authenticated using Auth0
func Auth0Middleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No authorization header",
			})
		}

		// Check if the header has the Bearer prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		// Get the token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Store token in context for later use
		c.Locals("token", token)

		return c.Next()
	}
}
