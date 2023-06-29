package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/damasdev/fiber/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:     os.Getenv("APP_NAME"),
		Prefork:     true,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Return status code with error message
			return c.Status(code).JSON(response.NewErrResponse(code, err.Error()))
		},
	}
}
