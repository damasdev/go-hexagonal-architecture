package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/damasdev/fiber/pkg/log"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	// Define server settings.
	isPrefork, _ := strconv.ParseBool(os.Getenv("SERVER_PREFORK"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName: os.Getenv("APP_NAME"),
		Prefork: isPrefork,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			message := err.Error()
			resp := response.New()

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code == fiber.StatusNotFound {
				message = response.StatusText(code)
			}

			resp.Status.Code = code
			resp.Status.Message = message

			switch code {
			case fiber.StatusInternalServerError:
				logger.Logger.Error(message, log.WithError(err))
			default:
				logger.Logger.Warning(message)
			}

			// Return status code with error message
			return c.Status(code).JSON(resp)
		},
	}
}
