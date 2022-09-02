package command

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type AddNewManual struct {
}

func (this *AddNewManual) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		log.Print("AddNewManual Error")
		return nil, fmt.Errorf("Dont have, telegram query")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.Text = "Введите название фирмы и модель к которой будет загружен мануал"
	server.Util.Stage[msg.From.ID] = server.Addmanualinfo
	return json.Marshal(Answer)
}
