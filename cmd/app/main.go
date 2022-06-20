package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://docker:mongopw@localhost:55001"))
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("Test").Collection("Users")
	var usr repository.User
	usr.TelegramID = 1111
	usr.UserID = "1"
	usr.Username = "debic"
	_, err = collection.InsertOne(context.TODO(), bson.D{{"name", "id"}, {"test", 2}})
	if err != nil {
		fmt.Println(err.Error())
	}

}
