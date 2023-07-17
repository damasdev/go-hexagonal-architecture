package user

import (
	port "go-hexagonal-architecture/internal/core/port/user"
	"go-hexagonal-architecture/internal/interfaces/http/response"
	"go-hexagonal-architecture/pkg/log"
	"go-hexagonal-architecture/pkg/logger"

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

	resp := response.New()
	resp.Status.Code = response.StatusOK
	resp.Status.Message = response.StatusText(response.StatusOK)
	resp.Data = "Hello World"

	return c.JSON(resp)
}
