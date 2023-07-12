package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go-hexagonal-architecture/internal/interfaces/http/routes"
	"go-hexagonal-architecture/pkg/config"
	"go-hexagonal-architecture/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.LoadEnvVars()
	config.ConnectMongoDB()
}

func main() {
	// Init Server
	server := fiber.New(config.FiberConfig())

	// Register Middleware
	server.Use(recover.New())
	server.Use(helmet.New())

	// Register Router
	routes.API(server)

	level, err := strconv.ParseInt(os.Getenv("LOG_THRESHOLD"), 10, 64)
	if err != nil {
		level = int64(logger.WarnLevel)
	}

	logger.Initialize(
		logger.WithWriter(os.Stdout),
		logger.WithLevel(logger.LogLevel(level)),
		logger.WithName(os.Getenv("APP_NAME")),
	)

	go func() {
		if err := server.Listen(":" + os.Getenv("APP_PORT")); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
