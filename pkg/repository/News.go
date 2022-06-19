package repository

import "time"

type News struct {
	NewsID     string    `bson:"newsid"`
	Text       string    `bson:"text"`
	AsRead     bool      `bson:"asread"`
	Time       time.Time `bson:"time"`
	ConsumerID int64     `bson:"consumerid"`
}
