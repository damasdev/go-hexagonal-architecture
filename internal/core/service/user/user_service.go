package user

import (
	port "github.com/damasdev/fiber/internal/core/port/user"
)

type service struct {
	repository port.UserRepository
}

func New(repository port.UserRepository) port.UserService {
	return &service{
		repository: repository,
	}
}
