package repository

import (
	"context"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Usermongo struct {
	*mongo.Collection
}

func NewUsermongo(collection *mongo.Collection) *Usermongo {
	return &Usermongo{
		collection,
	}
}

func (m *Usermongo) CreateUser(TelegramID int64, username string) error {
	var usr pubrep.User
	usr.TelegramID = TelegramID
	opt := options.Update().SetUpsert(true)
	filter := bson.D{{"telegram_id", TelegramID}}

	usr.Username = username
	data, err := bson.Marshal(usr)
	if err != nil {
		return fmt.Errorf("Bson Marshal err")
	}
	_, err = m.UpdateOne(context.TODO(), filter, data, opt)
	if err != nil {
		return fmt.Errorf("Insert Error")
	}
	return nil
}
func (m *Usermongo) UpdateUser(NewData pubrep.User) error {

}
func (m *Usermongo) GetUser(TelegramID int64) ([]pubrep.User, error) {

}
func (m *Usermongo) DeleteUser(TeelgramID int64) error {

}
