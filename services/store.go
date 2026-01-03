package services

import (
	"encoding/json"

	"github.com/abdulmalikraji/e-commerce/db/dao/storeDao"
	"github.com/abdulmalikraji/e-commerce/db/dao/userDao"
	"github.com/abdulmalikraji/e-commerce/db/models"
	"github.com/abdulmalikraji/e-commerce/dto/storeDto"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/auth-go"
)

type StoreService interface {
	CreateStore(ctx *fiber.Ctx, request storeDto.CreateStoreRequest) (int, error)
	GetStoreByID(ctx *fiber.Ctx, storeID string) (*models.Store, error)
	FindStore(ctx *fiber.Ctx, name string) ([]models.Store, error)
}

type storeService struct {
	userDao    userDao.DataAccess
	authClient auth.Client
	storeDao   storeDao.DataAccess
}

func NewStoreService(
	userDao userDao.DataAccess,
	authClient auth.Client,
	storeDao storeDao.DataAccess,
) StoreService {
	return storeService{
		userDao:    userDao,
		authClient: authClient,
		storeDao:   storeDao,
	}
}

func (s storeService) CreateStore(ctx *fiber.Ctx, request storeDto.CreateStoreRequest) (int, error) {
	// Verify Owner exists
	owner, err := s.userDao.FindById(request.OwnerID)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	// Create Store
	settingsBytes, err := json.Marshal(request.Settings)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	storeSettings := string(settingsBytes)

	_, err = s.storeDao.Insert(models.Store{
		Name:        request.Name,
		Description: request.Description,
		OwnerID:     owner.ID,
		Settings:    storeSettings,
	})
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

//todo: implement get store by id
func (s storeService) GetStoreByID(ctx *fiber.Ctx, storeID string) (*models.Store, error) {
	return nil, nil
}

//todo: implement find store
func (s storeService) FindStore(ctx *fiber.Ctx, name string) ([]models.Store, error) {
	return nil, nil
}