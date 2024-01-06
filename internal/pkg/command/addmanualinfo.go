package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type AddManualInfo struct {
	Firm     pubrep.Firms
	FirmChan chan pubrep.Firm
}

func (this *AddManualInfo) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.Message)
	if !ok {
		log.Println("AddManualInfo handl dont have TelegramMEssage in input parametr")
		return nil, fmt.Errorf("Не содержит сообщения. Попробуйте отправить заного.")
	}
	var Answer types.SendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	ManualData := strings.Split(msg.Text, " ")

	if len(ManualData) != 2 {
		log.Print("Входящее сообщение не содержит необходимого количества парметров")
		Answer.Text = "Отправьте в сообщении только название фирмы и модели разделенные пробелом"
		return json.Marshal(Answer)
	}
	var Manual pubrep.Manual
	ManualFirm := ManualData[0]
	DeviceModel := ManualData[1]
	rslt, err := this.Firm.GetFirm(ManualFirm)
	if err != nil {
		Answer.Text = "Произошла внутренняя ошибка. Попробуйте начать сначала или обратитесь к администратору"
		delete(server.Util.Stage, msg.From.ID)
		return util.CommandErrorHandler(&Answer, err)
	}
	if len(rslt) == 0 {
		FirmId, err := this.Firm.CreateFirm(ManualFirm)
		if err != nil {
			delete(server.Util.Stage, msg.From.ID)
			Answer.Text = "Произошла внутреняя ошибка"
			return util.CommandErrorHandler(&Answer, err)
		}

		var NewFirm pubrep.Firm
		NewFirm.ID = FirmId
		NewFirm.Firm = ManualFirm
		this.FirmChan <- NewFirm
		Manual.Firm = NewFirm
	} else {
		Manual.Firm = rslt[0]
	}

	Manual.DeviceModel = DeviceModel
	server.Util.Manual[msg.From.ID] = Manual
	server.Util.Stage[msg.From.ID] = server.AddManualDocument
	Answer.Text = "Отправьте в этот чат файл с мануалом"
	return json.Marshal(Answer)

}
