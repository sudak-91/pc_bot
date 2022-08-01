package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type News struct {
	NewsID     uuid.UUID `bson:"_id"`
	Text       string    `bson:"text"`
	AsRead     bool      `bson:"asread"`
	Time       time.Time `bson:"time"`
	ConsumerID int64     `bson:"consumerid"`
}
