package user

import (
	"github.com/damasdev/fiber/pkg/log"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	defer func() {
		logger.Logger.Debug("hello world", log.WithData(map[string]any{"hello": "world"}))
	}()
	return c.JSON("Hello, World!")
}
