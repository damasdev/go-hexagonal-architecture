package user

import (
	port "github.com/damasdev/fiber/internal/core/port/user"
)

type service struct {
	userRepo port.UserRepository
}

func New(userRepository port.UserRepository) port.UserService {
	return &service{
		userRepo: userRepository,
	}
}
