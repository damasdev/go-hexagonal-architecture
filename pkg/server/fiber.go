package server

import (
	"log"
	"os"
	"os/signal"
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

	logger.Initialize(
		logger.WithLevel(logger.InfoLevel),
		logger.WithName("fiber"),
	)

	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
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
