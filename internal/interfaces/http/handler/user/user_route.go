package user

import (
	"github.com/damasdev/fiber/pkg/log"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/response"
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

	response := response.DefaultResponse{}
	response.Status.Code = fiber.StatusOK
	response.Status.Message = "OK"
	response.Data = "Hello World"

	return c.JSON(response)
}
