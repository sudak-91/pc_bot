package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct {
	ID    primitive.ObjectID `bson:"_id"`
	Model string             `bson:"model"`
}
