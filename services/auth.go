package services

import (
	"time"

	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/db/dao/userTokenDao"
	"github.com/abdulmalikraji/e-commerce/db/models"
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
	LoginByEmail(ctx *fiber.Ctx, request authDto.LoginByEmailRequest) (authDto.LoginByEmailResponse, int, error)
	ValidateToken(ctx *fiber.Ctx, token string) error
	RefreshToken(ctx *fiber.Ctx) (authDto.RefreshTokenResponse, error)
	Logout(ctx *fiber.Ctx, request authDto.LogoutRequest) error
}

type authService struct {
	userDao      userDao.DataAccess
	authClient   auth.Client
	userTokenDao userTokenDao.DataAccess
}

func NewAuthService(
	userDao userDao.DataAccess,
	authClient auth.Client,
	userTokenDao userTokenDao.DataAccess,
) AuthService {
	return authService{
		userDao:      userDao,
		authClient:   authClient,
		userTokenDao: userTokenDao,
	}
}

func (s authService) LoginByEmail(ctx *fiber.Ctx, request authDto.LoginByEmailRequest) (authDto.LoginByEmailResponse, int, error) {
	resp, err := s.authClient.Token(types.TokenRequest{
		GrantType: "password",
		Email:     request.Email,
		Password:  request.Password,
	})

	if err != nil {
		return authDto.LoginByEmailResponse{}, fiber.StatusUnauthorized, err
	}

	// Set refresh token in HTTP-only cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    resp.RefreshToken,
		Path:     "/",
		Expires:  time.Unix(resp.ExpiresAt, 0),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Strict",
	})

	// Store the user token in the database
	_, err = s.userTokenDao.Insert(models.UserToken{
		UserID:       resp.User.ID,
		RefreshToken: resp.RefreshToken,
		IsRevoked:    false,
		ExpiresAt:    time.Unix(resp.ExpiresAt, 0),
	})
	if err != nil {
		return authDto.LoginByEmailResponse{}, fiber.StatusInternalServerError, err
	}

	// Set the Authorization header
	ctx.Set("Authorization", "Bearer "+resp.AccessToken)

	// Return data for client storage (localStorage)
	return authDto.LoginByEmailResponse{
		AccessToken: resp.AccessToken,
		ExpiresAt:   time.Unix(resp.ExpiresAt, 0),
	}, fiber.StatusOK, nil
}

func (s authService) ValidateToken(ctx *fiber.Ctx, token string) error {
	_, err := s.authClient.WithToken(token).GetUser()
	if err != nil {
		return err
	}
	return nil
}

func (s authService) RefreshToken(ctx *fiber.Ctx) (authDto.RefreshTokenResponse, error) {
	// Get refresh token from cookie
	refreshToken := ctx.Cookies("refresh_token")
	if refreshToken == "" {
		return authDto.RefreshTokenResponse{}, fiber.NewError(fiber.StatusUnauthorized, "No refresh token found")
	}

	// Check if the token is revoked
	token, err := s.userTokenDao.FindByRefreshToken(refreshToken)
	if err != nil {
		return authDto.RefreshTokenResponse{}, err
	}
	if token.IsRevoked {
		return authDto.RefreshTokenResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Refresh token has been revoked")
	}

	//revoke the old refresh token
	err = s.userTokenDao.RevokeToken(token.ID.String())
	if err != nil {
		return authDto.RefreshTokenResponse{}, err
	}

	// Get new access token
	resp, err := s.authClient.RefreshToken(refreshToken)
	if err != nil {
		return authDto.RefreshTokenResponse{}, err
	}

	// Set the Authorization header
	ctx.Set("Authorization", "Bearer "+resp.AccessToken)

	//update the refresh token cookie
	// Set refresh token in HTTP-only cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    resp.RefreshToken,
		Path:     "/",
		Expires:  time.Unix(resp.ExpiresAt, 0),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Strict",
	})

	// Store the user token in the database
	_, err = s.userTokenDao.Insert(models.UserToken{
		UserID:       resp.User.ID,
		RefreshToken: resp.RefreshToken,
		IsRevoked:    false,
		ExpiresAt:    time.Unix(resp.ExpiresAt, 0),
	})

	return authDto.RefreshTokenResponse{
		AccessToken: resp.AccessToken,
		ExpiresAt:   time.Unix(resp.ExpiresAt, 0),
	}, nil
}

func (s authService) Logout(ctx *fiber.Ctx, request authDto.LogoutRequest) error {
	// Invalidate the access token
	err := s.authClient.WithToken(request.AccessToken).Logout()
	if err != nil {
		return err
	}
	// Clear the refresh token cookie
	ctx.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})
	return nil
}
