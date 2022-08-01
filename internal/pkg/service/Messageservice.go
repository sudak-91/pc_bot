package service

import (
	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

func (t *TelegramUpdater) messageService(Message types.TelegramMessage) ([]byte, error) {
	switch server.Util.Stage[Message.From.ID] {
	case 0:
		//Старт обработки обычной текстовой команды
		return t.Routing(Message)
	case 10:
		return t.Execute("/addnews", Message)

	default:
		return t.DefaultAnswer(&Message)

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
