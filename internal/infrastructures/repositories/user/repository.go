package user

import (
	port "github.com/damasdev/go-hexagonal-architecture/internal/core/ports/user"
)

// repository is a struct that represents the user repository
type repository struct{}

// New creates a new instance of the Repository
func New() port.Repository {
	return &repository{}
}
