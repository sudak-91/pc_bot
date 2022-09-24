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
	"github.com/sudak-91/pc_bot/internal/pkg/notificator"
	intrep "github.com/sudak-91/pc_bot/internal/pkg/repository"
	"github.com/sudak-91/pc_bot/internal/pkg/server"
	intserv "github.com/sudak-91/pc_bot/internal/pkg/service"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
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
	//Make keyboard
	mkeyboard := createMainInlineKeyboard()
	akeyboard := createAdminInlineKeyboard()
	FirmChan := make(chan pubrep.Firm, 3)
	ManualChan := make(chan pubrep.Manual, 3)

	NotificationService := notificator.NewNotification(ManualChan, FirmChan)
	go NotificationService.Run()

	addBotCommand(telegramupdate, repo, mkeyboard, akeyboard, FirmChan, ManualChan)

	BotServer := server.NewServer(viper.GetString("server.port"), os.Getenv("BOT_KEY"), updater)
	AdminUsr, err := repo.Users.GetAdmin()
	var AdminID int64
	if len(AdminUsr) != 0 {
		AdminID = AdminUsr[0].TelegramID
	}
	if err != nil {
		log.Println("No admin")
		AdminID = 0
	}
	BotServer.Run(AdminID)
}

//addBotCommand -  adding command handler
func addBotCommand(telegramupdate *intserv.TelegramUpdater, repo *intrep.MongoRepository, mkeyboard keyboardmaker.InlineCommandKeyboard, akeyboard keyboardmaker.InlineCommandKeyboard,
	FirmChan chan pubrep.Firm, ManualChan chan pubrep.Manual) {
	telegramupdate.AddNewCommand("/default", &intcom.Default{})
	telegramupdate.AddNewCommand("/start", &intcom.StartCommand{User: repo.Users, Keyboard: mkeyboard.GetKeyboard()})
	telegramupdate.AddNewCommand("/login", &intcom.Login{Users: repo.Users, Keyboard: akeyboard.GetKeyboard()})
	telegramupdate.AddNewCommand("/news", &intcom.News{})
	telegramupdate.AddNewCommand("/addnews", &intcom.AddNews{News: repo.Newser})
	telegramupdate.AddNewCommand("/question", &intcom.Questions{})
	telegramupdate.AddNewCommand("/addquestion", &intcom.AddQuestion{Question: repo.Questions})
	telegramupdate.AddNewCommand("/shown", &intcom.Shown{News: repo.Newser})
	telegramupdate.AddNewCommand("/readmore", &intcom.ReadMore{News: repo.Newser})
	telegramupdate.AddNewCommand("/markasread", &intcom.MarkAsRead{News: repo.Newser})
	telegramupdate.AddNewCommand("/showq", &intcom.ShowQ{Question: repo.Questions})
	telegramupdate.AddNewCommand("/sendanswer", &intcom.SendAnswer{Questions: repo.Questions})
	telegramupdate.AddNewCommand("/sendanswerto", &intcom.SendAnswerTo{Question: repo.Questions})
	telegramupdate.AddNewCommand("/markasanswer", &intcom.MarkAsAnswer{Question: repo.Questions})
	telegramupdate.AddNewCommand("/addmanual", &intcom.AddNewManual{})
	telegramupdate.AddNewCommand("/addmanualinfo", &intcom.AddManualInfo{Firm: repo.Firm, FirmChan: FirmChan})
	telegramupdate.AddNewCommand("/addmanualdocument", &intcom.AddManualDocument{Manual: repo.Manual})
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
	mainkeyboard.AddButton("Задать вопрос", "/question", 0, 0)
	mainkeyboard.AddButton("Предложить новость", "/news", 0, 1)
	return mainkeyboard
}

func createAdminInlineKeyboard() keyboardmaker.InlineCommandKeyboard {
	var adminkeyboard keyboardmaker.InlineCommandKeyboard
	adminkeyboard.MakeGrid(2, 2)
	adminkeyboard.AddButton("Показать все вопросы", "/showq", 0, 0)
	adminkeyboard.AddButton("Show all news", "/shown", 0, 1)
	adminkeyboard.AddButton("Добавить мануал", "/addmanual", 1, 0)
	return adminkeyboard
}
