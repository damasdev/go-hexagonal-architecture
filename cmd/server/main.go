package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"go-hexagonal-architecture/internal/interfaces/http/server"
	"go-hexagonal-architecture/pkg/config"
)

func init() {
	config.LoadEnvVars()
	config.ConnectMongoDB()
}

func main() {
	app := server.NewFiberServer()

	app.RegisterMiddleware()
	app.RegisterHandler()

	go func() {
		if err := app.Run(); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Stop(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
