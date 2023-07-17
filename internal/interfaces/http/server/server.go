package server

import (
	"context"
	"os"
	"strconv"

	handler "go-hexagonal-architecture/internal/interfaces/http/handler/errors"
	"go-hexagonal-architecture/internal/interfaces/http/routes"
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
		ErrorHandler: handler.ErrorHandler,
	}
}
