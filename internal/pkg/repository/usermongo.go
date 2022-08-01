package repository

import (
	"context"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Usermongo struct {
	col *mongo.Collection
}

func NewUsermongo(db *mongo.Database) *Usermongo {
	var u Usermongo
	u.col = db.Collection("Users")

	//defer u.col.Drop(context.TODO())
	return &u
}

func (m *Usermongo) CreateUser(TelegramID int64, username string) error {
	var usr pubrep.User
	usr.TelegramID = TelegramID
	usr.Username = username
	usr.Role = 0

	data, err := bson.Marshal(usr)
	if err != nil {
		return fmt.Errorf("Bson Marshal err")
	}
	_, err = m.col.InsertOne(context.TODO(), data)

	if !mongo.IsDuplicateKeyError(err) {
		return fmt.Errorf("Insert Error: %s", err.Error())
	}
	return nil
}
func (m *Usermongo) UpdateUser(NewData pubrep.User) error {
	filter := bson.D{{"_id", NewData.TelegramID}}
	upd := bson.D{{"$set", bson.D{{"Username", NewData.Username}, {"Role", NewData.Role}}}}
	_, err := m.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return fmt.Errorf("Update user has error: %s", err.Error())
	}
	return nil
}

func (m *Usermongo) GetUser(TelegramID int64) ([]pubrep.User, error) {
	filter := bson.D{{"_id", TelegramID}}
	rtslt := m.col.FindOne(context.TODO(), filter)
	Users := make([]pubrep.User, 1)
	err := rtslt.Decode(&Users[0])
	if err != nil {
		return nil, fmt.Errorf("GetUser has error: %s", err.Error())
	}
	return Users, nil
}

func (m *Usermongo) GetAdmin() ([]pubrep.User, error) {
	filter := bson.D{{"Role", 9}}
	rtslt := m.col.FindOne(context.TODO(), filter)
	Users := make([]pubrep.User, 1)
	err := rtslt.Decode(&Users[0])
	if err != nil {
		return nil, fmt.Errorf("GetUser has error: %s", err.Error())
	}
	return Users, nil
}
func (m *Usermongo) GetUsers() ([]pubrep.User, error) {
	rslt, err := m.col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("Get users has error: %s", err.Error())
	}
	var us []pubrep.User
	err = rslt.All(context.TODO(), &us)
	if err != nil {
		return nil, fmt.Errorf("Get users has error elem: %s", err.Error())
	}

	return us, nil

}

func (m *Usermongo) DeleteUser(TelegramID int64) error {
	m.col.DeleteOne(context.TODO(), bson.D{{"_id", TelegramID}})
	return nil
}

func (m *Usermongo) DeleteAll() error {
	_, err := m.col.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return fmt.Errorf("DeleteAll has error: %s", err.Error())
	}
	return nil
}
