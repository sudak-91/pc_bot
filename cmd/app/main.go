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

	akeyboard := createAdminInlineKeyboard()
	FirmChan := make(chan pubrep.Firm, 3)
	ManualChan := make(chan pubrep.Manual, 3)

	NotificationService := notificator.NewNotification(ManualChan, FirmChan)
	go NotificationService.Run()

	addBotCommand(telegramupdate, repo, akeyboard, FirmChan, ManualChan)

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
func addBotCommand(telegramupdate *intserv.TelegramUpdater, repo *intrep.MongoRepository, akeyboard keyboardmaker.InlineCommandKeyboard,
	FirmChan chan pubrep.Firm, ManualChan chan pubrep.Manual) {
	telegramupdate.AddNewCommand("/default", &intcom.Default{})
	telegramupdate.AddNewCommand("/start", &intcom.StartCommand{User: repo.Users})
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
	telegramupdate.AddNewCommand("/addmanualdocument", &intcom.AddManualDocument{Manual: repo.Manual, ManualNotificator: ManualChan})
	telegramupdate.AddNewCommand("/editfirm", &intcom.EditFirmCommand{Firms: repo.Firm})
	telegramupdate.AddNewCommand("/confirmeditfirm", &intcom.ConfirmEditFirm{Firms: repo.Firm, Manuals: repo.Manual})
	telegramupdate.AddNewCommand("/approvedfirm", &intcom.ApprovedFirm{Firms: repo.Firm})
	telegramupdate.AddNewCommand("/editmanual", &intcom.EditManual{Manuals: repo.Manual})
	telegramupdate.AddNewCommand("/confirmeditmanual", &intcom.ConfirmEditManual{Manual: repo.Manual})
	telegramupdate.AddNewCommand("/approvedmanual", &intcom.ApprovedManual{Manual: repo.Manual})
	telegramupdate.AddNewCommand("/manualarchive", &intcom.ManualArchive{})
	telegramupdate.AddNewCommand("/allfirmslist", &intcom.AllFirmsList{Firms: repo.Firm})
	telegramupdate.AddNewCommand("/allunapprovedfirmslist", &intcom.AllUnapprovedFirmsList{Firms: repo.Firm})
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
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
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

func createAdminInlineKeyboard() keyboardmaker.InlineCommandKeyboard {
	var adminkeyboard keyboardmaker.InlineCommandKeyboard
	adminkeyboard.MakeGrid(3, 3)
	adminkeyboard.AddButton("Показать все вопросы", "/showq", 0, 0)
	adminkeyboard.AddButton("Show all news", "/shown", 0, 1)
	adminkeyboard.AddButton("Добавить мануал", "/addmanual", 0, 2)
	adminkeyboard.AddButton("Показать мануалы", "/allfirmslist 0", 1, 0)
	adminkeyboard.AddButton("Unapproved manuals", "/allunapprovedfirmslist 0", 1, 1)
	return adminkeyboard
}
