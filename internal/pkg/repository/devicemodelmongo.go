package repository

import (
	"context"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceModelMongo struct {
	col *mongo.Collection
}

func NewDeviceModelMongo(db *mongo.Database) *DeviceModelMongo {
	var d DeviceModelMongo
	d.col = db.Collection("DeviceModel")
	return &d
}

func (f *DeviceModelMongo) CreateModel(DeviceName string) (primitive.ObjectID, error) {
	var NewDeviceModel pubrep.DeviceModel
	NewDeviceModel.Model = DeviceName
	data, err := bson.Marshal(NewDeviceModel)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", err.Error())
	}
	rslt, err := f.col.InsertOne(context.TODO(), data)
	if !mongo.IsDuplicateKeyError(err) {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", err.Error())
	}
	retrslt, ok := rslt.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", "Not object id")
	}
	return retrslt, nil
}

func (f *DeviceModelMongo) UpdateModel(NewModel pubrep.DeviceModel) error {
	filter := bson.D{{"_id", NewModel.Model}}
	upd := bson.D{{"$set", bson.D{{"approved", NewModel.Approved}}}}
	_, err := f.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("UpdateModel has error: %s", err.Error())
	}
	return nil
}

func (f *DeviceModelMongo) GetModel(Name string) ([]pubrep.DeviceModel, error) {
	filter := bson.D{{"_id", fmt.Sprintf("/%s/", Name)}}

	cursor, err := f.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())
	}
	var Result []pubrep.DeviceModel
	err = cursor.All(context.TODO(), &Result)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())

	}
	return Result, nil
}
func (f *DeviceModelMongo) DeleteModel(Name string) error {
	//TODO: Add delete firm logic
	return nil
}
