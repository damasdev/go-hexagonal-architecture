package user

import (
	port "github.com/damasdev/fiber/internal/core/port/user"
	"github.com/damasdev/fiber/pkg/log"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service port.UserService
}

func New(service port.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(c *fiber.Ctx) error {
	defer func() {
		logger.Logger.Debug("hello world", log.WithData(map[string]any{"hello": "world"}))
	}()

	json := response.New()
	json.Status.Code = fiber.StatusOK
	json.Status.Message = response.StatusOK
	json.Data = "Hello World"

	return c.JSON(json)
}
