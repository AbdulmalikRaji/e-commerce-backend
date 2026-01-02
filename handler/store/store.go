package store

import (
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/gofiber/fiber/v2"
)

type StoreHandler interface {
	CreateStore(ctx *fiber.Ctx) error
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
	return nil
}
