package user

import (
	port "github.com/damasdev/fiber/internal/core/port/user"
)

type repository struct{}

func New() port.UserRepository {
	return &repository{}
}
