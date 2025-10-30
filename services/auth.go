package services

import (
	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

type AuthService interface {
	LoginByEmail(ctx *fiber.Ctx, loginRequest authDto.LoginByEmailRequest) (string, error)
	ValidateToken(ctx *fiber.Ctx, token string) (bool, error)
	RefreshToken(ctx *fiber.Ctx, oldToken string) (string, error)
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

func (s authService) LoginByEmail(ctx *fiber.Ctx, loginRequest authDto.LoginByEmailRequest) (string, error) {
	resp, err := s.authClient.Token(types.TokenRequest{
		GrantType: "password",
		Email:     loginRequest.Email,
		Password:  loginRequest.Password,
	})

	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

func (s authService) ValidateToken(ctx *fiber.Ctx, token string) (bool, error) {
	return true, nil
}

func (s authService) RefreshToken(ctx *fiber.Ctx, oldToken string) (string, error) {

	return "", nil
}
