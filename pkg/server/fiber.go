package server

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	routes "github.com/damasdev/fiber/internal/interfaces/http"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type fiberServer struct{}

func NewFiber() Server {
	return &fiberServer{}
}

func (server *fiberServer) Run() {
	app := fiber.New()

	app.Use(recover.New())

	routes.API(app)

	level, err := strconv.ParseInt(os.Getenv("LOG_THRESHOLD"), 10, 64)
	if err != nil {
		level = int64(logger.InfoLevel)
	}

	logger.Initialize(
		logger.WithLevel(logger.LogLevel(level)),
		logger.WithName(os.Getenv("APP_NAME")),
	)

	go func() {
		if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal(err.Error())
	}
}
