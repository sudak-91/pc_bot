package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

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
	TitleFirmName := strings.ToTitle(FirmName)
	var NewFirm pubrep.Firm
	NewFirm.Firm = TitleFirmName
	data, err := bson.Marshal(NewFirm)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("CreateFirm has error: %s", err.Error())
	}
	rslt, err := f.col.InsertOne(context.TODO(), data)

	if err != nil {
		log.Println(err.Error())
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
	upd := bson.D{{"$set", bson.D{{"approved", NewFirm.Approved}, {"firm", NewFirm.Firm}}}}
	_, err := f.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("UpdateModel has error: %s", err.Error())
	}
	return nil
}

func (f *FirmsMongo) GetFirm(Name string) ([]pubrep.Firm, error) {
	TitleFirmName := strings.ToTitle(Name)
	filter := bson.D{{"firm", TitleFirmName}}

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

//GetFirms return all filtred firms
func (f *FirmsMongo) GetFirms(filter string) ([]pubrep.Firm, error) {
	cursore, err := f.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetFirms has error: %s", err.Error())
	}
	var Result []pubrep.Firm
	err = cursore.All(context.TODO(), &Result)
	if err != nil {
		return nil, fmt.Errorf("GetFirms has error: %s", err.Error())
	}
	return Result, nil
}

func (f *FirmsMongo) GetFirmById(ID string) ([]pubrep.Firm, error) {
	log.Printf("Input ID is: %s", ID)
	ObjID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("GetFitmById has error:%s\n", err.Error())
	}
	filter := bson.D{{"_id", ObjID}}

	rslt := f.col.FindOne(context.TODO(), filter)
	var FirmFromID pubrep.Firm
	err = rslt.Decode(&FirmFromID)
	if err != nil {
		return nil, fmt.Errorf("GetModel has error: %s", err.Error())

	}
	var Result []pubrep.Firm
	Result = append(Result, FirmFromID)
	return Result, nil
}
func (f *FirmsMongo) DeleteFirm(Name string) error {
	//TODO: Add delete firm logic
	return nil
}
