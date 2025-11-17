package services

import (
	"os"
	"time"

	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/db/dao/userTokenDao"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
	"gorm.io/gorm"
)

const (
	accessTokenKey  = "access_token"
	refreshTokenKey = "refresh_token"
	expiresAtKey    = "expires_at"
)

type AuthService interface {
	LoginByEmail(ctx *fiber.Ctx, request authDto.LoginByEmailRequest) (authDto.LoginByEmailResponse, int, error)
	SignupByEmail(ctx *fiber.Ctx, request authDto.SignUpByEmailRequest) (int, error)
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
		log.Errorf("login failed for email=%s: %v", request.Email, err)
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
		log.Errorf("failed to store user token for user_id=%s: %v", resp.User.ID, err)
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

func (s authService) SignupByEmail(ctx *fiber.Ctx, request authDto.SignUpByEmailRequest) (int, error) {

	if s.userDao.IsEmailExists(request.Email) {
		return fiber.StatusConflict, fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	// Create user in Auth0
	userResp, err := s.authClient.Signup(types.SignupRequest{
		Email:    request.Email,
		Password: request.Password,
		Data: map[string]interface{}{
			"first_name": request.Firstname,
			"last_name":  request.Lastname,
			"role":       request.Role,
		},
	})
	if err != nil && userResp.ID == uuid.Nil {
		return fiber.StatusInternalServerError, err
	}

	// Store user in the database inside a transaction so DB changes are rolled back on error.
	err = s.userDao.Transaction(func(tx *gorm.DB) error {
		newUser := models.User{
			Auth0ID:   userResp.ID,
			Email:     request.Email,
			FirstName: request.Firstname,
			LastName:  request.Lastname,
		}

		res := tx.Table(newUser.TableName()).Create(&newUser)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if err != nil {
		// Attempt best-effort cleanup: delete the auth provider user we just created
		if delErr := s.authClient.WithToken(os.Getenv("SERVICE_ROLE_KEY")).AdminDeleteUser(types.AdminDeleteUserRequest{UserID: userResp.ID}); delErr != nil {
			log.Warnf("cleanup failed for auth user %s after DB error: %v", userResp.ID.String(), delErr)
		} else {
			log.Infof("cleaned up auth user %s after DB error", userResp.ID.String())
		}

		log.Errorf("signup failed for email=%s: %v", request.Email, err)
		return fiber.StatusInternalServerError, err
	}
	// Audit log - do not log sensitive data
	log.Infof("signup success user_id=%s email=%s", userResp.ID.String(), request.Email)
	// Post-creation: start background goroutine to enqueue verification email/welcome job.
	go func(uid string, email string) {
		// Placeholder background task — replace with real queueing (Redis/RabbitMQ) later.
		log.Infof("background job placeholder: send verification email for user_id=%s email=%s", uid, email)
	}(userResp.ID.String(), request.Email)

	log.Infof("enqueued post-signup tasks for user_id=%s", userResp.ID.String())

	return fiber.StatusCreated, nil
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

	token := ctx.Get("Authorization")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing authorization header")
	}
	// Invalidate the access token
	err := s.authClient.WithToken(token).Logout()
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
