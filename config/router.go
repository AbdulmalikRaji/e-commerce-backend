package config

import (
	"github.com/abdulmalikraji/e-commerce/config/middleware"
	"github.com/abdulmalikraji/e-commerce/db/connection"
	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/handler/authentication"
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/auth-go"
)

func InitializeRoutes(app *fiber.App, client connection.Client, auth auth.Client) {
	// Initialize services
	userDao := userDao.New(client)
	authService := services.NewAuthService(userDao, auth)
	authHandler := authentication.New(authService)

	// Create auth middleware
	tokenMiddleware := middleware.TokenValidationMiddleware(authService)

	// Auth routes (no token required)
	authGroup := app.Group("/auth")
	authGroup.Post("/login", authHandler.LoginByEmail)
	authGroup.Post("/refresh", authHandler.RefreshToken)
	authGroup.Post("/signup", authHandler.SignUp)
	authGroup.Get("/validate", tokenMiddleware, authHandler.ValidateToken)

	// Protected routes (require valid token)
	app.Use(tokenMiddleware) // Apply to all routes after this point
	authGroup.Post("/logout", authHandler.Logout)
}
