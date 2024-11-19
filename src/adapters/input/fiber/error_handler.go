package input_fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ErrorHandler func(*fiber.Ctx, error) error

func handleNotFoundError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).SendString(err.Error())
}

func handleConflictError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).SendString(err.Error())
}

func handleInternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
}

var errorHandlers = map[string]ErrorHandler{
	"*main.NotFoundError": handleNotFoundError,
	"*main.ConflictError": handleConflictError,
	"default":             handleInternalServerError,
}

func HandleError(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	errorType := fmt.Sprintf("%T", err)

	if handler, exists := errorHandlers[errorType]; exists {
		return handler(c, err)
	}

	return errorHandlers["default"](c, err)
}
