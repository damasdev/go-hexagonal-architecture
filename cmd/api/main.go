package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/damasdev/fiber/pkg/config"
	"github.com/damasdev/fiber/pkg/logger"
	"github.com/damasdev/fiber/pkg/server"
	"github.com/rs/zerolog/diode"
)

func main() {

	server := server.New()

	server.RegisterMiddleware()
	server.RegisterHandler()

	go func() {
		if err := server.Run(); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	config.LoadEnvVars()

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
}
