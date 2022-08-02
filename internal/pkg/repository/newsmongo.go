package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
)

type NewsMongo struct {
	col *mongo.Collection
}

func NewNewsMongo(db *mongo.Database) *NewsMongo {
	return &NewsMongo{
		col: db.Collection("News"),
	}
}

func (n *NewsMongo) CreateNews(Text string, ConsumerID int64) error {
	var NewNews pubrep.News
	NewNews.AsRead = false
	NewNews.ConsumerID = ConsumerID
	NewNews.Text = Text
	NewNews.Time = time.Now()
	id, err := uuid.New()
	if err != nil {
		return fmt.Errorf("Create News has error: %s", err.Error())
	}
	NewNews.NewsID = id
	data, err := bson.Marshal(NewNews)
	if err != nil {
		return fmt.Errorf("CreateNews has marshal error: %s", err.Error())
	}
	_, err = n.col.InsertOne(context.TODO(), data)
	if err != nil {
		return fmt.Errorf("CreateNews has InsertOne error: %s", err.Error())
	}
	return nil
}
func (n *NewsMongo) GetAllNews() ([]pubrep.News, error) {
	rslt, err := n.col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("GetAllNews has error^ %s", err.Error())
	}
	var NewsList []pubrep.News
	err = rslt.All(context.TODO(), &NewsList)
	if err != nil {
		return nil, err
	}
	return NewsList, nil
}
func (n *NewsMongo) GetNotAsReadNews() ([]pubrep.News, error) {
	rslt, err := n.col.Find(context.TODO(), bson.D{{"asread", false}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.News
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}

func (n *NewsMongo) GetAsReadNews() ([]pubrep.News, error) {
	rslt, err := n.col.Find(context.TODO(), bson.D{{"asread", true}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.News
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}

func (n *NewsMongo) GetNewsWithDate(time time.Time) ([]pubrep.News, error) {
	rslt, err := n.col.Find(context.TODO(), bson.D{{"time", time}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.News
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}

func (n *NewsMongo) GetNews(UUID string) ([]pubrep.News, error) {
	//FIXME: разобратся с выдаче UUID
	log.Printf("GetNews has UUID: %x\n", UUID)

	filter := bson.D{{"_id", fmt.Sprintf("%x", UUID)}}
	rtslt := n.col.FindOne(context.TODO(), filter)
	News := make([]pubrep.News, 1)
	err := rtslt.Decode(&News[0])
	if err != nil {
		return nil, fmt.Errorf("GetNews has error: %s", err.Error())
	}
	return News, nil
}

func (n *NewsMongo) GetNewsFromConsumer(ConsumerID int64) ([]pubrep.News, error) {
	rslt, err := n.col.Find(context.TODO(), bson.D{{"consumerid", ConsumerID}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.News
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}

func (n *NewsMongo) UpdateNews(NewNews pubrep.News) error {
	filter := bson.D{{"_id", NewNews.NewsID}}
	upd := bson.D{{"$set", bson.D{{"asread", NewNews.AsRead}, {"text", NewNews.Text}}}}
	_, err := n.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("Update user has error: %s", err.Error())
	}
	return nil
}
func (n *NewsMongo) DeleteNews(NewsID uuid.UUID) error {
	return nil
}
