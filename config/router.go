package config

import (
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/handler/authentication"
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/auth-go"
)

func InitializeRoutes(app *fiber.App, client connection.Client, auth auth.Client) {
	// Initialize session store

	// Initialize services
	userDao := userDao.New(client)
	authService := services.NewAuthService(userDao, auth)
	authHandler := authentication.New(authService)

	// Middleware to make authenticator available in context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("authenticator", auth)
		return c.Next()
	})

	// Auth routes
	authGroup := app.Group("/auth")
	authGroup.Post("/login", authHandler.LoginByEmail)
}
