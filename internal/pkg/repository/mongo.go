package repository

import (
	"github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Users     repository.Users
	Newser    repository.Newser
	Questions repository.Questions
	Client    *mongo.Client
}

func NewMongoRepository(UserRepo repository.Users, NewsRepo repository.Newser, QuestionRepo repository.Questions, client *mongo.Client) *MongoRepository {
	return &MongoRepository{
		Users:     NewUsermongo(),
		Newser:    NewsRepo,
		Questions: QuestionRepo,
		Client:    client,
	}
}
