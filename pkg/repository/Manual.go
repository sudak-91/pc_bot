package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manual struct {
	ManualID    primitive.ObjectID `bson:"_id"`
	FileUniqID  string             `bson:"fileid"`
	FirmName    string             `bson:"firm"`
	DeviceModel string             `bson:"model"`
	Version     string             `bson:"version,omitempty"`
	Approved    bool               `bson:"approved"`
}
