package repository

import (
	"context"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FirmsMongo struct {
	col *mongo.Collection
}

func NewFirmsMongo(db *mongo.Database) *FirmsMongo {
	return &FirmsMongo{
		col: db.Collection("Firms"),
	}
}

func (f *FirmsMongo) CreateFirm(FirmName string) (primitive.ObjectID, error) {
	var NewFirm pubrep.Firm
	NewFirm.Firm = FirmName
	data, err := bson.Marshal(NewFirm)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", err.Error())
	}
	rslt, err := f.col.InsertOne(context.TODO(), data)

	if !mongo.IsDuplicateKeyError(err) {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", err.Error())
	}
	retrslt, ok := rslt.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", "return not ID")
	}
	return retrslt, nil
}

func (f *FirmsMongo) UpdateFirm(NewFirm pubrep.Firm) error {
	filter := bson.D{{"_id", NewFirm.ID}}
	upd := bson.D{{"$set", bson.D{{"approved", NewFirm.Approved}}}}
	_, err := f.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("UpdateModel has error: %s", err.Error())
	}
	return nil
}

func (f *FirmsMongo) GetFirm(Name string) ([]pubrep.Firm, error) {
	filter := bson.D{{"firm", fmt.Sprintf("/%s/", Name)}}

	cursor, err := f.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())
	}
	var Result []pubrep.Firm
	err = cursor.All(context.TODO(), &Result)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())

	}
	return Result, nil
}
func (f *FirmsMongo) DeleteFirm(Name string) error {
	//TODO: Add delete firm logic
	return nil
}
