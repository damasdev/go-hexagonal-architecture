package server

import (
	"context"
	"errors"
	"os"
	"strconv"

	"go-hexagonal-architecture/internal/interfaces/http/response"
	"go-hexagonal-architecture/internal/interfaces/http/routes"
	"go-hexagonal-architecture/pkg/log"
	"go-hexagonal-architecture/pkg/logger"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server interface {
	Run() error
	Stop(ctx context.Context) error

	RegisterMiddleware()
	RegisterHandler()
}

type fiberServer struct {
	*fiber.App
}

func NewFiberServer() Server {
	server := fiber.New(LoadFiberConfig())

	level, err := strconv.ParseInt(os.Getenv("LOG_THRESHOLD"), 10, 64)
	if err != nil {
		level = int64(logger.WarnLevel)
	}

	logger.Initialize(
		logger.WithWriter(os.Stdout),
		logger.WithLevel(logger.LogLevel(level)),
		logger.WithName(os.Getenv("APP_NAME")),
	)

	return &fiberServer{
		server,
	}
}

func (f *fiberServer) Run() error {
	return f.Listen(":" + os.Getenv("APP_PORT"))
}

func (f *fiberServer) Stop(ctx context.Context) error {
	return f.ShutdownWithContext(ctx)
}

func (f *fiberServer) RegisterMiddleware() {
	f.Use(recover.New())
	f.Use(helmet.New())
}

func (f *fiberServer) RegisterHandler() {
	routes.API(f.App)
}

func LoadFiberConfig() fiber.Config {
	// Define server settings.
	isPrefork, _ := strconv.ParseBool(os.Getenv("SERVER_PREFORK"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:      os.Getenv("APP_NAME"),
		Prefork:      isPrefork,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: errorHandler,
	}
}

var errorHandler = func(c *fiber.Ctx, err error) error {
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
}
