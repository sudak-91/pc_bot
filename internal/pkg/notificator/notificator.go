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
	device chan pubrep.DeviceModel
}

func NewNotification(Manual chan pubrep.Manual, Firm chan pubrep.Firm, Device chan pubrep.DeviceModel) *Notification {
	return &Notification{
		manual: Manual,
		firm:   Firm,
		device: Device,
	}
}

func (n *Notification) Run() {
	for {
		select {
		case manual := <-n.manual:
			go sendManualNotification(manual)
		case firm := <-n.firm:
			go sendAddFirmNotification(firm)
		case device := <-n.device:
			go addDeviceNotofocation(device)
		}
	}
}

func sendManualNotification(manual pubrep.Manual) {
	var message methods.SendMessage
	message.Text = fmt.Sprintf("Получен новый мануал\n Фирма: %s\n Модель:%s\n", manual.FirmName, manual.DeviceModel)
	message.ChatID = server.Util.AdminID
	err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), message)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s : %d has error %s", file, line, err.Error())
	}

}

func sendAddFirmNotification(firm pubrep.Firm) {

}

func addDeviceNotofocation(device pubrep.DeviceModel) {

}
