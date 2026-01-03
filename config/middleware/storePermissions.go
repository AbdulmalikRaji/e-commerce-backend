package middleware

import (
	"github.com/abdulmalikraji/e-commerce/db/dao/storeUserDao"
	"github.com/abdulmalikraji/e-commerce/utils/genericResponse"
	"github.com/abdulmalikraji/e-commerce/utils/messages"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// StorePermissionMiddleware returns a middleware that checks whether the
// requesting user (read from `X-User-ID` header) has the given action
// permission on the store identified by route param `store_id`.
//
// Example usage: dao := storeUserDao.New(connClient); app.Use("/stores/:store_id/orders", StorePermissionMiddleware(dao, storeUserDao.ActionManageOrders))
// StorePermissionMiddleware checks whether the requesting user (whose ID
// must already be set into the request context with `c.Locals("user_id")`)
// has the given action permission on the store identified by route param `store_id`.
func StorePermissionMiddleware(dao storeUserDao.DataAccess, action string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get store id from common param names
		storeParam := c.Params("store_id")
		if storeParam == "" {
			return genericResponse.ErrorResponse(c, fiber.StatusBadRequest, "missing store id in path")
		}

		storeID, err := uuid.Parse(storeParam)
		if err != nil {
			return genericResponse.ErrorResponse(c, fiber.StatusBadRequest, "invalid store id format")
		}

		// Read the user id header set during login: `X-User-ID`.
		uid := c.Get("X-User-ID")
		if uid == "" {
			return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized, messages.CreateMsg(c, messages.Unauthorized))
		}
		userID, err := uuid.Parse(uid)
		if err != nil {
			return genericResponse.ErrorResponse(c, fiber.StatusUnauthorized, messages.CreateMsg(c, messages.Unauthorized))
		}

		if !dao.HasPermission(storeID, userID, action) {
			return genericResponse.ErrorResponse(c, fiber.StatusForbidden, messages.CreateMsg(c, messages.Unauthorized))
		}

		return c.Next()
	}
}
