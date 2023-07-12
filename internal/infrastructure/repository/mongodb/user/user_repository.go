package user

import (
	port "go-hexagonal-architecture/internal/core/port/user"

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
