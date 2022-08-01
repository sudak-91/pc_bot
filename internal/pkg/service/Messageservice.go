package service

import (
	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

//TODO: Обработка ошибок
func (t *TelegramUpdater) messageService(Message types.TelegramMessage) ([]byte, error) {
	switch server.Util.Stage[Message.From.ID] {

	case 10:
		return t.Execute("/addnews", Message)
	case 20:
		return t.Execute("/addquestion", Message)
	default:
		return t.Routing(Message)

	}
}

func (t *TelegramUpdater) Routing(Message types.TelegramMessage) ([]byte, error) {
	for _, ent := range Message.Entities {
		switch ent.Type {
		case "bot_command":
			return t.Execute(Message.Text, Message)
		default:
			return nil, nil
		}
	}
	return nil, nil
}
