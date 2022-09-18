package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Firm struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Firm     string             `bson:"firm"`
	Approved bool               `bson:"approved,omitempty"`
}
