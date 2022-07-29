package command

import (
	"encoding/json"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type StartCommand struct {
	User     pubrep.Users
	Keyboard types.TelegramInlineKeyboardMarkup
}

func (s *StartCommand) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("It is not a message")
	}
	err := s.User.CreateUser(msg.Chat.ID, msg.From.Username)
	if err != nil {
		return nil, fmt.Errorf("StartCommand handler has error: %s", err.Error())
	}
	var Answer types.TelegramSendMessage
	Answer.ChatID = msg.Chat.ID
	Answer.Method = "sendMessage"
	var Linktochannel types.TelegramMessageEntity
	Linktochannel.Length = 24
	Linktochannel.Type = "text_link"
	Linktochannel.Url = "https://t.me/wtfcontrolsengineer"
	Linktochannel.Offset = 40
	Answer.Entities = append(Answer.Entities, Linktochannel)
	Answer.Text = `Добро пожаловать. Бот создан для канала \"Я вам че-Автоматизатор\"`
	Answer.ReplyMarkup = &s.Keyboard
	return json.Marshal(Answer)
}
