package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	intrep "github.com/sudak-91/pc_bot/internal/pkg/repository"
	"github.com/sudak-91/pc_bot/internal/pkg/server"
	intserv "github.com/sudak-91/pc_bot/internal/pkg/service"
	update "github.com/sudak-91/telegrambotgo/Service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := initConf()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load()
	db := createMongoClientAndPing()
	repo := intrep.NewMongoRepository(db)
	//Создание стандартного обработчика обновления
	//telegramupdater экеземпляр, который содержит реализацию обработок всех основных обновлдений
	telegramupdate := intserv.NewTelegramUpdater()
	//updater - роутинг для обновлений
	updater := update.NewTelegramService(telegramupdate)

	BotServer := server.NewServer(viper.GetString("server.port"), os.Getenv("BOT_KEY"), updater)
	BotServer.Run()
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

func initConf() error {
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

}
