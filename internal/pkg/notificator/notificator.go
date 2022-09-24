package notificator

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	methods "github.com/sudak-91/telegrambotgo/TelegramAPI/Methods"
)

type Notification struct {
	manual chan pubrep.Manual
	firm   chan pubrep.Firm
}

func NewNotification(Manual chan pubrep.Manual, Firm chan pubrep.Firm) *Notification {
	return &Notification{
		manual: Manual,
		firm:   Firm,
	}
}

func (n *Notification) Run() {
	for {
		select {
		case manual := <-n.manual:
			go sendManualNotification(manual)
		case firm := <-n.firm:
			go sendAddFirmNotification(firm)

		}
	}
}

func sendManualNotification(manual pubrep.Manual) {
	var message methods.SendMessage
	//FIXME: требуется вытащить значения фирмы и девайса из документов ибо сейчас там только ID
	message.Text = fmt.Sprintf("Получен новый мануал\n Фирма: %s\n Модель:%s\n", manual.FirmName, manual.DeviceModel)
	message.ChatID = server.Util.AdminID
	err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), message)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s : %d has error %s", file, line, err.Error())
	}
	var doc methods.SendDocument
	doc.Document = manual.FileUniqID
	doc.ChatId = server.Util.AdminID
	err = methods.SendDocumentMethod(os.Getenv("BOT_KEY"), doc)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s : %d has error %s", file, line, err.Error())
	}
}

//TODO: Добаивть клавиатуру с возможностью редактирования или подтверждения
func sendAddFirmNotification(firm pubrep.Firm) {
	var message methods.SendMessage
	message.Text = fmt.Sprintf("Добавлена новая фирма: %s\n", firm.Firm)
	message.ChatID = server.Util.AdminID
	err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), message)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d has error %s", file, line, err.Error())
	}
}
