package repository

type Firm struct {
	Firm     string `bson:"_id"`
	Approved bool   `bson:"approved,omitempty"`
}