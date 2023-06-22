package server

import "github.com/damasdev/fiber/pkg/logger"

type Server interface {
	Run()
}

type server struct {
	logger logger.Logger
}

func New(logger logger.Logger) Server {
	return &server{
		logger: logger,
	}
}

func (srv *server) Run() {
	srv.logger.Info("server is running..")
}
