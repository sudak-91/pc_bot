package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type DeviceModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Model    string             `bson:"model"`
	Approved bool               `bson:"approved,omitempty"`
}
