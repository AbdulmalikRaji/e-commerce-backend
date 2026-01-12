package genericResponse

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)

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

	// Attempt to extract a JSON object from the error message string
	trimmed := strings.TrimSpace(msg)
	// Look for a JSON object inside the message (e.g. "response status code 400: {...}")
	jsonStart := strings.Index(trimmed, "{")
	jsonEnd := strings.LastIndex(trimmed, "}")
	if jsonStart != -1 && jsonEnd != -1 && jsonEnd > jsonStart {
		jsonStr := trimmed[jsonStart : jsonEnd+1]
		var parsed map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &parsed); err == nil {
			// If caller didn't provide explicit data, expose the parsed JSON as Data
			if len(data) == 0 {
				out = parsed
			}

			// Prefer common message keys from parsed JSON to form a clean Message
			if m, ok := parsed["msg"].(string); ok && m != "" {
				trimmed = m
			} else if m, ok := parsed["message"].(string); ok && m != "" {
				trimmed = m
			} else if m, ok := parsed["error"].(string); ok && m != "" {
				trimmed = m
			} else if m, ok := parsed["error_description"].(string); ok && m != "" {
				trimmed = m
			}
		}
	}

	return ctx.Status(status).JSON(GenericResponse{
		Success: false,
		Message: trimmed,
		Data:    out,
	})
}
