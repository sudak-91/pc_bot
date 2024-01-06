package repository

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	QuestionID    uuid.UUID `bson:"_id"`
	Text          string    `bson:"text"`
	Date          time.Time `bson:"date"`
	AsAnswer      bool      `bson:"asanswer"`
	ContributerID int64     `bson:"contributerid"`
	MessageID     int64     `bson:"messageid"`
}
