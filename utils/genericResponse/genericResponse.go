package genericResponse

import "github.com/gofiber/fiber/v2"

type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *fiber.Ctx, status int, data interface{}, msg ...string) error {
	if len(msg) == 0 {
		msg = append(msg, "Success")
	}

	return ctx.Status(status).JSON(GenericResponse{
		Success: true,
		Message: msg[0],
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status int, msg string, data ...interface{}) error {
	// If caller provided exactly one data argument, return it directly as Data.
	// Otherwise return the slice of provided data (or empty slice if none).
	var out interface{}
	if len(data) == 0 {
		out = []interface{}{}
	} else if len(data) == 1 {
		out = data[0]
	} else {
		out = data
	}

	return ctx.Status(status).JSON(GenericResponse{
		Success: false,
		Message: msg,
		Data:    out,
	})
}
