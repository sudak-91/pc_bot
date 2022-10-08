package command

import (
	"encoding/json"
	"fmt"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type StartCommand struct {
	User pubrep.Users
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
	Answer.Text = `Добро пожаловать. Бот создан для канала "Я вам че-Автоматизатор"`
	Answer.ReplyMarkup = s.createKeyboard()
	return json.Marshal(Answer)
}

func (s *StartCommand) createKeyboard() *types.TelegramInlineKeyboardMarkup {
	var mainkeyboard keyboardmaker.InlineCommandKeyboard
	mainkeyboard.MakeGrid(1, 3)
	mainkeyboard.AddButton("Задать вопрос", "/question", 0, 0)
	mainkeyboard.AddButton("Предложить новость", "/news", 0, 1)
	mainkeyboard.AddButton("Архив мануалов", "/manualarchive", 0, 2)
	rslt := mainkeyboard.GetKeyboard()
	return &rslt
}
