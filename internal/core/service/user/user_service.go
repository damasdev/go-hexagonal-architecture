package user

import (
	port "go-hexagonal-architecture/internal/core/port/user"
)

type service struct {
	userRepo port.UserRepository
}

func New(userRepository port.UserRepository) port.UserService {
	return &service{
		userRepo: userRepository,
	}
}
