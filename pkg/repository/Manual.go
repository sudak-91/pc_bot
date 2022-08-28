package repository

type Manual struct {
	FileUniqID  string `bson:"_id"`
	FirmName    string `bson:"firmid"`
	DeviceModel string `bson:"modelid"`
	Version     string `bson:"version"`
	Approved    bool   `bson:"approved"`
}
