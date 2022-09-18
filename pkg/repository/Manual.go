package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manual struct {
	FileUniqID  string             `bson:"_id"`
	FirmName    primitive.ObjectID `bson:"firmid"`
	DeviceModel primitive.ObjectID `bson:"modelid"`
	Version     string             `bson:"version,omitempty"`
	Approved    bool               `bson:"approved"`
}
