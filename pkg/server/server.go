package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	routes "github.com/damasdev/fiber/internal/interfaces/http"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/diode"
)

type Server interface {
	Run()
}

type fiberServer struct{}

func New() Server {
	return &fiberServer{}
}

func (server *fiberServer) Run() {
	app := fiber.New()

	app.Use(recover.New())

	level, err := strconv.ParseInt(os.Getenv("LOG_THRESHOLD"), 10, 64)
	if err != nil {
		level = int64(logger.WarnLevel)
	}

	logger.Initialize(
		logger.WithWriter(
			diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
				fmt.Printf("Logger Dropped %d messages", missed)
			}),
		),
		logger.WithLevel(logger.LogLevel(level)),
		logger.WithName(os.Getenv("APP_NAME")),
	)

	routes.API(app)

	go func() {
		if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Fatal(err.Error())
	}
}
