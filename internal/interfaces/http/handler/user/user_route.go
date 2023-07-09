package user

import (
	port "github.com/damasdev/fiber/internal/core/port/user"
	"github.com/damasdev/fiber/pkg/log"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService port.UserService
}

func New(userService port.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) List(c *fiber.Ctx) error {
	defer func() {
		logger.Logger.Debug("hello world", log.WithData(map[string]any{"hello": "world"}))
	}()

	json := response.New()
	json.Status.Code = response.StatusOK
	json.Status.Message = response.StatusText(response.StatusOK)
	json.Data = "Hello World"

	return c.JSON(json)
}
