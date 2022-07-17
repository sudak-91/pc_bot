package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Question struct {
	QuestionID    uuid.UUID `bson:"_id"`
	Text          string    `bson:"text"`
	Date          time.Time `bson:"date"`
	AsAnswer      bool      `bson:"asanswer"`
	ContributerID int64     `bson:"contributerid"`
}
