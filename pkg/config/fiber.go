package config

import (
	"errors"
	"os"
	"strconv"

	"go-hexagonal-architecture/pkg/log"
	"go-hexagonal-architecture/pkg/logger"
	"go-hexagonal-architecture/pkg/response"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	// Define server settings.
	isPrefork, _ := strconv.ParseBool(os.Getenv("SERVER_PREFORK"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:     os.Getenv("APP_NAME"),
		Prefork:     isPrefork,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
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
				logger.Logger.Error(message, log.WithError(err), log.WithSkip(3))
			default:
				logger.Logger.Warning(message, log.WithError(err), log.WithSkip(3))
			}

			// Return status code with error message
			return c.Status(code).JSON(resp)
		},
	}
}
