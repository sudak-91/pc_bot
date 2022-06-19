package repository

import "time"

type Question struct {
	QuestionID string    `bson:"questionid"`
	Text       string    `bson:"text"`
	Date       time.Time `bson:"date"`
	AsAnswer   bool      `bson:"asanswer"`
	ConsumerID int64     `bson:"consumerid"`
}
