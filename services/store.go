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
	GetStoreByID(ctx *fiber.Ctx, request storeDto.GetStoreByIDRequest) (storeDto.GetStoreByIDResponse, int, error)
	FindStore(ctx *fiber.Ctx, request storeDto.FindStoreRequest) (storeDto.FindStoreResponse, int, error)
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

func (s storeService) GetStoreByID(ctx *fiber.Ctx, request storeDto.GetStoreByIDRequest) (storeDto.GetStoreByIDResponse, int, error) {
	store, err := s.storeDao.FindById(request.StoreID)
	if err != nil {
		return storeDto.GetStoreByIDResponse{}, fiber.StatusInternalServerError, err
	}

	var storeImage string
	if store.Image != nil {
		storeImage = *store.Image
	}

	return storeDto.GetStoreByIDResponse{
		ID:          store.ID.String(),
		Name:        store.Name,
		Description: store.Description,
		OwnerID:     store.OwnerID.String(),
		Image:       storeImage,
		Settings:    store.Settings,
	}, fiber.StatusOK, nil
}

func (s storeService) FindStore(ctx *fiber.Ctx, request storeDto.FindStoreRequest) (storeDto.FindStoreResponse, int, error) {
	stores, err := s.storeDao.FindByName(request.Name)
	if err != nil {
		return storeDto.FindStoreResponse{}, fiber.StatusInternalServerError, err
	}

	var storeSummaries []storeDto.StoreSummary
	for _, store := range stores {
		var storeImage string
		if store.Image != nil {
			storeImage = *store.Image
		}
		storeSummaries = append(storeSummaries, storeDto.StoreSummary{
			ID:          store.ID.String(),
			Name:        store.Name,
			Image:       storeImage,
			Description: store.Description,
			OwnerID:     store.OwnerID.String(),
		})
	}

	return storeDto.FindStoreResponse{
		Stores: storeSummaries,
	}, fiber.StatusOK, nil
}
