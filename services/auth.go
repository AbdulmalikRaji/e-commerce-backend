package services

import (
	"time"

	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

const (
	accessTokenKey  = "access_token"
	refreshTokenKey = "refresh_token"
	expiresAtKey    = "expires_at"
)

type AuthService interface {
	LoginByEmail(ctx *fiber.Ctx, loginRequest authDto.LoginByEmailRequest) (authDto.LoginByEmailResponse, int, error)
	ValidateToken(ctx *fiber.Ctx, token string) (bool, error)
	RefreshToken(ctx *fiber.Ctx, refreshToken string) (string, error)
	GetAccessToken(ctx *fiber.Ctx) string
	GetRefreshToken(ctx *fiber.Ctx) string
	SetTokens(ctx *fiber.Ctx, accessToken, refreshToken string, expiresAt int64)
	IsTokenExpired(ctx *fiber.Ctx) bool
}

type authService struct {
	userDao    userDao.DataAccess
	authClient auth.Client
}

func NewAuthService(userDao userDao.DataAccess, authClient auth.Client) AuthService {
	return authService{
		userDao:    userDao,
		authClient: authClient,
	}
}

func (s authService) LoginByEmail(ctx *fiber.Ctx, loginRequest authDto.LoginByEmailRequest) (authDto.LoginByEmailResponse, int, error) {
	resp, err := s.authClient.Token(types.TokenRequest{
		GrantType: "password",
		Email:     loginRequest.Email,
		Password:  loginRequest.Password,
	})

	if err != nil {
		return authDto.LoginByEmailResponse{}, fiber.StatusUnauthorized, err
	}

	// Set the Authorization header with the new token
	ctx.Set("Authorization", "Bearer "+resp.AccessToken)

	// Store tokens in context
	s.SetTokens(ctx, resp.AccessToken, resp.RefreshToken, resp.ExpiresAt)

	return authDto.LoginByEmailResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		ExpiresAt:    resp.ExpiresAt,
	}, fiber.StatusOK, nil
}

// GetAccessToken retrieves the access token from the context
func (s authService) GetAccessToken(ctx *fiber.Ctx) string {
	return ctx.Locals(accessTokenKey).(string)
}

// GetRefreshToken retrieves the refresh token from the context
func (s authService) GetRefreshToken(ctx *fiber.Ctx) string {
	return ctx.Locals(refreshTokenKey).(string)
}

// SetTokens stores the tokens and expiry in the context
func (s authService) SetTokens(ctx *fiber.Ctx, accessToken, refreshToken string, expiresAt int64) {
	ctx.Locals(accessTokenKey, accessToken)
	ctx.Locals(refreshTokenKey, refreshToken)
	ctx.Locals(expiresAtKey, expiresAt)
}

// IsTokenExpired checks if the current access token is expired
func (s authService) IsTokenExpired(ctx *fiber.Ctx) bool {
	expiresAt, ok := ctx.Locals(expiresAtKey).(int64)
	if !ok {
		return true
	}
	return time.Now().Unix() >= expiresAt
}

func (s authService) ValidateToken(ctx *fiber.Ctx, token string) (bool, error) {
	_, err := s.authClient.WithToken(token).GetUser()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s authService) RefreshToken(ctx *fiber.Ctx, refreshToken string) (string, error) {
	resp, err := s.authClient.RefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
