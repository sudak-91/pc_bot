package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	intcom "github.com/sudak-91/pc_bot/internal/pkg/command"
	intrep "github.com/sudak-91/pc_bot/internal/pkg/repository"
	"github.com/sudak-91/pc_bot/internal/pkg/server"
	intserv "github.com/sudak-91/pc_bot/internal/pkg/service"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	update "github.com/sudak-91/telegrambotgo/Service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	fmt.Println("Engcore bot started")
	fmt.Println("Engcore bot load config file")
	err := initConf()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Init complite")
	if viper.GetBool("dev") {
		err = godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}

	db := createMongoClientAndPing()
	repo := intrep.NewMongoRepository(db)
	//Создание стандартного обработчика обновления
	//telegramupdater экеземпляр, который содержит реализацию обработок всех основных обновлдений
	telegramupdate := intserv.NewTelegramUpdater()
	//updater - роутинг для обновлений
	updater := update.NewTelegramService(telegramupdate)
	//Сбор бизнес логики
	mkeyboard := createMainInlineKeyboard()

	telegramupdate.AddNewCommand("/start", &intcom.StartCommand{User: repo.Users, Keyboard: mkeyboard.GetKeyboard()})
	BotServer := server.NewServer(viper.GetString("server.port"), os.Getenv("BOT_KEY"), updater)
	BotServer.Run()
}

func createMongoClientAndPing() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	connectString := fmt.Sprintf("mongodb://%s:%s@mongodb:27017", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err != nil {
		panic(err.Error())
	}
	log.Println("connect complete")
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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func createMainInlineKeyboard() keyboardmaker.InlineCommandKeyboard {
	var mainkeyboard keyboardmaker.InlineCommandKeyboard
	mainkeyboard.MakeGrid(1, 2)
	mainkeyboard.AddButton("Задать вопрос", "/addquestion", 0, 0)
	mainkeyboard.AddButton("Предложить новость", "/addnews", 0, 1)
	return mainkeyboard
}
