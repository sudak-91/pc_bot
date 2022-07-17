package command

import (
	"encoding/json"
	"fmt"

	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type StartCommand struct {
}

func (s *StartCommand) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("It is not a message")
	}
	var Answer types.TelegramSendMessage
	Answer.ChatID = msg.Chat.ID
	Answer.Method = "sendMessage"
	Answer.Text = "Thx for you Request"
	return json.Marshal(Answer)
}
