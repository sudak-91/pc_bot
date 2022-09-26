package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manual struct {
	ManualID    primitive.ObjectID `bson:"_id"`
	FileUniqID  string             `bson:"fileid"`
	Firm        *Firm              `bson:"firm"`
	DeviceModel string             `bson:"device"`
	Contributer int64              `bosn:"userid"`
	Version     string             `bson:"version,omitempty"`
	Approved    bool               `bson:"approved"`
}
