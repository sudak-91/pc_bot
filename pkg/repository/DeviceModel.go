package repository

type DeviceModel struct {
	Model    string `bson:"_id"`
	Approved bool   `bson:"approved,omitempty"`
}
