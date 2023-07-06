package user

import (
	port "github.com/damasdev/fiber/internal/core/port/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func New(db *mongo.Database) port.UserRepository {
	return &repository{
		db: db,
	}
}
