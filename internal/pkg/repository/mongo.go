package repository

import (
	"github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Users     repository.Users
	Newser    repository.Newser
	Questions repository.Questions
	Firm      repository.Firms
	Manual    repository.Manuals
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Users:     NewUsermongo(db),
		Newser:    NewNewsMongo(db),
		Questions: NewQuestionmongo(db),
		Firm:      NewFirmsMongo(db),
		Manual:    NewManualMong(db),
	}
}
