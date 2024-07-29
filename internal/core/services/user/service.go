package user

import (
	"errors"

	port "github.com/damasdev/go-hexagonal-architecture/internal/core/ports/user"
)

// Config is a struct that represents the configuration of the service
type Config struct {
	// Repository is a port.Repository instance
	Repository port.Repository
}

// Validate validates the Config
func (c Config) Validate() error {
	if c.Repository == nil {
		return errors.New("missing repository")
	}
	return nil
}

// service is a struct that represents the user service
type service struct {
	// repository is a port.Repository instance
	repository port.Repository
}

// New creates a new instance of the Service
func New(cfg Config) port.Service {
	if err := cfg.Validate(); err != nil {
		return nil
	}

	return &service{
		repository: cfg.Repository,
	}
}
