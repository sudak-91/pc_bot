package repository

import (
	"context"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManualMongo struct {
	col *mongo.Collection
}

func NewManualMong(db *mongo.Database) *ManualMongo {
	var m ManualMongo
	m.col = db.Collection("Manuals")
	return &m
}

func (m *ManualMongo) CreateManual(NewManual pubrep.Manual) error {
	data, err := bson.Marshal(NewManual)
	if err != nil {
		return fmt.Errorf("CreateManual has error: %s", err.Error())
	}
	_, err = m.col.InsertOne(context.TODO(), data)
	if err != nil {
		return fmt.Errorf("CreateFirm has error: %s", err.Error())
	}
	return nil
}

func (m *ManualMongo) UpdateManual(NewManual pubrep.Manual) error {
	filter := bson.D{{"_id", NewManual.FileUniqID}}
	updData, err := bson.Marshal(NewManual)
	if err != nil {
		return fmt.Errorf("UpdateManual has error: %s", err.Error())
	}

	upd := bson.D{{"$set", updData}}
	_, err = m.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("UpdateModel has error: %s", err.Error())
	}
	return nil

}

//db.Manuals.aggregate({$lookup: {from: "Firms", localField: "firmid", foreignField: "_id", as: "firmname" }  } )
func (m *ManualMongo) GetManual(Firm string, DeviceName string) ([]pubrep.Manual, error) {

	filter := bson.D{{"firm", Firm}, {"device", DeviceName}}
	cursor, err := m.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())
	}
	var Result []pubrep.Manual
	err = cursor.All(context.TODO(), &Result)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())

	}
	return Result, nil
}

func (m *ManualMongo) DeleteModel(ID string) error {
	//TODO: Add Delete Model Logic
	return nil
}
