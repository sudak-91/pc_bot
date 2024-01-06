package command

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type AddNewManual struct {
}

// AddNewManual handl is a entry point to machine state of adding manual to db
func (this *AddNewManual) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.CallbackQuery)
	if !ok {
		log.Print("AddNewManual Error")
		return nil, fmt.Errorf("Dont have, telegram query")
	}
	var Answer types.SendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	Answer.Text = "Введите название фирмы и модель к которой будет загружен мануал"
	server.Util.Stage[msg.From.ID] = server.AddManualInfo
	return json.Marshal(Answer)
}
