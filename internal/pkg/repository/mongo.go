package repository

import (
	"github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Users     repository.Users
	Newser    repository.Newser
	Questions repository.Questions
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Users: NewUsermongo(db),
	}
}
