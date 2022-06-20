package repository

import (
	"github.com/sudak-91/pc_bot/pkg/repository"
)

type MongoRepository struct {
	repository.Users
	repository.Newser
	repository.Questions
}

func NewMongoRepository(UserRepo repository.Users, NewsRepo repository.Newser, QuestionRepo repository.Questions) *MongoRepository {
	return &MongoRepository{
		UserRepo,
		NewsRepo,
		QuestionRepo,
	}
}
