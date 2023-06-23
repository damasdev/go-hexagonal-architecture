package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	routes "github.com/damasdev/fiber/internal/interfaces/http"
	"github.com/gofiber/fiber/v2"
)

type fiberServer struct{}

func NewFiber() Server {
	return &fiberServer{}
}

func (server *fiberServer) Run() {
	app := fiber.New()

	routes.API(app)

	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
