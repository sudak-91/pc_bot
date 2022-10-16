package repository

import (
	"context"
	"fmt"
	"log"
	"runtime"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s : %d has error %s", file, line, err.Error())
		return fmt.Errorf("CreateManual has error: %s", err.Error())
	}
	_, err = m.col.InsertOne(context.TODO(), data)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s : %d has error %s", file, line, err.Error())
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

func (m *ManualMongo) UpdateEmbeddedFirm(NewFirm pubrep.Firm) error {
	log.Println(NewFirm)
	filter := bson.D{{"firm._id", NewFirm.ID}}
	upd := bson.D{{"$set",
		bson.D{
			{"firm.firm", NewFirm.Firm}}}}

	rslt, err := m.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("UodateEmbeddedFirm has error: %s", err.Error())
	}
	log.Println(rslt.MatchedCount, rslt.ModifiedCount)
	return nil
}

//db.Manuals.aggregate({$lookup: {from: "Firms", localField: "firmid", foreignField: "_id", as: "firmname" }  } )
func (m *ManualMongo) GetManuals() ([]pubrep.Manual, error) {
	cursor, err := m.col.Find(context.TODO(), bson.D{})
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

func (m *ManualMongo) GetManualByID(ID string) (pubrep.Manual, error) {
	ObjID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return pubrep.Manual{}, fmt.Errorf("GetManualByID has error: %s", err.Error())

	}
	filter := bson.D{{"_id", ObjID}}
	rslt := m.col.FindOne(context.TODO(), filter)
	var Manual pubrep.Manual
	err = rslt.Decode(&Manual)
	if err != nil {
		return Manual, fmt.Errorf("GetManualByID has error: %s", err.Error())
	}
	return Manual, nil
}
func (m *ManualMongo) GetApprovedManuals(approved bool) ([]pubrep.Manual, error) {
	filter := bson.D{{"approved", approved}}
	cursor, err := m.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var Manuals []pubrep.Manual
	err = cursor.All(context.TODO(), &Manuals)
	if err != nil {
		return nil, err
	}
	return Manuals, nil

}

func (m *ManualMongo) DeleteManuals(ID string) error {
	//TODO: Add Delete Model Logic
	return nil
}
