package command

import (
	"encoding/json"
	"fmt"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type Questions struct {
}

func (q *Questions) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.CallbackQuery)
	var Answer types.SendMessage
	if !ok {
		return nil, fmt.Errorf("Questions handl dont have Callbackquery type on the input parametr")
	}
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	Answer.Text = "Можете указать ваш вопрос далее"
	server.Util.Stage[msg.From.ID] = 20
	return json.Marshal(Answer)
}
