package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Firm struct {
	ID   primitive.ObjectID `bson:"id"`
	Firm string             `bson:"firm"`
}
