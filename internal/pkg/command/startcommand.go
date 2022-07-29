package command

import (
	"encoding/json"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type StartCommand struct {
	pubrep.Users
}

func (s *StartCommand) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("It is not a message")
	}
	var Answer types.TelegramSendMessage
	Answer.ChatID = msg.Chat.ID
	Answer.Method = "sendMessage"
	var Linktochannel types.TelegramMessageEntity
	Linktochannel.Length = 24
	Linktochannel.Type = "url"
	Linktochannel.Url = "https://t.me/wtfcontrolsengineer"
	Linktochannel.Offset = 41
	Answer.Entities = append(Answer.Entities, Linktochannel)
	Answer.Text = "Добро пожаловать. Бот создан для канала \"Я вам че-Автоматизатор\""
	return json.Marshal(Answer)
}
