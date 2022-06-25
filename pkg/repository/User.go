package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     string             `bson:"userid"`
	TelegramID int64              `bson:"telegram_id"`
	Username   string             `bson:"Username"`
}
