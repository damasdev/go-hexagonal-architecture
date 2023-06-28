package server

import (
	"errors"

	"github.com/damasdev/fiber/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// Default error handler
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Return status code with error message
	return c.Status(code).JSON(response.NewErrResponse(code, err.Error()))
}
