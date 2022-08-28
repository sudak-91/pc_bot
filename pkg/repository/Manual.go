package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Manual struct {
	ID         primitive.ObjectID `bson:"_id"`
	FirmID     string             `bson:"firmid"`
	ModelID    string             `bson:"modelid"`
	FileUniqID string             `bson:"file_uniq_id"`
	Version    string             `bson:"version"`
	Approved   bool               `bson:"approved"`
}
