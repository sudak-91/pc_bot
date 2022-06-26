package main

import (
	"context"
	"fmt"
	"time"

	intrep "github.com/sudak-91/pc_bot/internal/pkg/repository"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	db := createMongoClientAndPing()

	repo := intrep.NewMongoRepository(db)
	err := repo.Users.CreateUser(int64(4555667776), "test1")
	if err != nil {
		panic(err.Error())
	}
	err = repo.Users.CreateUser(int64(9934848941), "test2")
	if err != nil {
		panic(err.Error())
	}
	err = repo.Users.CreateUser(int64(4555667776), "test11")
	if err != nil {
		var newdata pubrep.User
		newdata.TelegramID = 4555667776
		newdata.Username = "test11"
		err2 := repo.Users.UpdateUser(newdata)
		if err2 != nil {
			panic(err2.Error())
		}
	}
	us, err := repo.Users.GetUser(int64(4555667776))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v", us[0].Username)
	us2, err := repo.Users.GetUsers()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v", us2)
	err = repo.Users.DeleteAll()
	if err != nil {
		panic(err.Error())
	}
}

func createMongoClientAndPing() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://docker:mongopw@localhost:55001"))
	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err != nil {
		panic(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err.Error())
	}
	db := client.Database("Test")
	return db
}
