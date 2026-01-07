package store

import (
	"github.com/abdulmalikraji/e-commerce/dto/storeDto"
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/gofiber/fiber/v2"
)

type StoreHandler interface {
	CreateStore(ctx *fiber.Ctx) error
	GetStoreByID(ctx *fiber.Ctx) error
}

type storeHandler struct {
	service services.StoreService
}

func New(service services.StoreService) StoreHandler {
	return storeHandler{
		service: service,
	}
}

func (c storeHandler) CreateStore(ctx *fiber.Ctx) error {
	var request storeDto.CreateStoreRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	statusCode, err := c.service.CreateStore(ctx, request)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return genericResponse.SuccessResponse(ctx, statusCode, nil, "Store created successfully")
}

func (c storeHandler) GetStoreByID(ctx *fiber.Ctx) error {
	var request storeDto.GetStoreByIDRequest
	if err := ctx.QueryParser(&request); err != nil {
		return err
	}

	response, statusCode, err := c.service.GetStoreByID(ctx, request)
	if err != nil {
		return genericResponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())

	}

	return genericResponse.SuccessResponse(ctx, statusCode, response, "Store retrieved successfully")
}
