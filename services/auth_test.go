package services

import (
	"testing"
	"time"

	authenticator "github.com/abdulmalikraji/e-commerce/mocks/auth"
	"github.com/abdulmalikraji/e-commerce/mocks/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/mocks/dao/userTokenDao"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/supabase-community/auth-go/types"
	"github.com/valyala/fasthttp"
)

var userMockDao *userDao.MockDataAccess
var userTokenMockDao *userTokenDao.MockDataAccess
var mockAuthClient *authenticator.MockClient

var fiberCtx *fiber.Ctx

var s AuthService

var user_id, _ = uuid.Parse("user_id")
var refreshExpiry = time.Now().Add(7 * 24 * time.Hour)

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	app := fiber.New()
	fiberCtx = app.AcquireCtx(&fasthttp.RequestCtx{})

	// Assign language to fiber context header
	fiberCtx.Request().Header.Set("Accept-Language", "en")

	userMockDao = userDao.NewMockDataAccess(ct)
	userTokenMockDao = userTokenDao.NewMockDataAccess(ct)
	mockAuthClient = authenticator.NewMockClient(ct)

	s = NewAuthService(userMockDao, mockAuthClient, userTokenMockDao)
	return func() {
		s = nil
		defer ct.Finish()
	}
}

func TestAuthService_Login_By_Email_Successfully(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockAuthClient.EXPECT().Token(types.TokenRequest{
		GrantType: "password",
		Email:     "jack.doe@company.com",
		Password:  "password",
	}).Return(types.TokenResponse{
		types.Session{
			User: types.User{
				ID: user_id,
			},
			RefreshToken: "refresh",
			AccessToken: "access"
		},
	})

	userTokenDao.EXPECT().Insert(models.UserToken{
		UserID:       user_id
		RefreshToken: "refresh",
		IsRevoked:    false,
		ExpiresAt:    refreshExpiry,
	})

	// userMockDao.EXPECT().FindByEmailOrPhoneNumber().Return(fakeUsers[2], nil)
	// cryptoMock.EXPECT().CheckPasswordHash("hashedPassword", fakeUsers[2].Password).Return(true)

	// response, status, err := s.Login(&fiber.Ctx{}, authDto.LoginRequest{
	// 	UniqueIdentifier: "jack.doe@company.com",
	// 	Password:         "hashedPassword",
	// })

	// if err != nil {
	// 	t.Errorf("Error when login")
	// }

	// assert.NotEmpty(t, response)
	// assert.Equal(t, fiber.StatusOK, status)
	// assert.Equal(t, response.FullName, fakeUsers[2].Name+" "+fakeUsers[2].Surname)
	// assert.Equal(t, response.Email, fakeUsers[2].Email)
	// assert.Equal(t, response.PhoneNumber, fakeUsers[2].PhoneNumber)
	// assert.NotNil(t, response.Token)
}
