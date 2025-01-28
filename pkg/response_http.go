package pkg

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Response404(c *fiber.Ctx) error {
	err := Response{
		Status:  fiber.StatusNotFound,
		Message: "Resource not found",
	}

	return ResponseJSON(c, fiber.StatusNotFound, err, nil)
}

func Response409(c *fiber.Ctx) error {
	err := Response{
		Status:  fiber.StatusConflict,
		Message: "Conflict occurred",
	}

	return ResponseJSON(c, fiber.StatusConflict, err, nil)
}

func Response500(c *fiber.Ctx) error {
	err := Response{
		Status:  fiber.StatusInternalServerError,
		Message: "Internal server error",
	}

	return ResponseJSON(c, fiber.StatusInternalServerError, err, nil)
}

func Response401(c *fiber.Ctx) error {
	err := Response{
		Status:  fiber.StatusUnauthorized,
		Message: "Unauthorized access",
	}

	return ResponseJSON(c, fiber.StatusUnauthorized, err, nil)
}

func ResponseJSON(c *fiber.Ctx, statusCode int, data any, headers map[string]string) error {
	response := Response{
		Status: statusCode,
		Data:   data,
	}

	for key, value := range headers {
		c.Set(key, value)
	}

	c.Set("Content-Type", "application/json")

	return c.Status(statusCode).JSON(response)
}
