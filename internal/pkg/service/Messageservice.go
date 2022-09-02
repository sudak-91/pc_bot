package service

import (
	"log"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

//TODO: Обработка ошибок
func (t *TelegramUpdater) messageService(Message types.TelegramMessage) ([]byte, error) {
	_, ok := server.Util.Stage[Message.From.ID]
	if !ok {
		return t.Routing(Message)
	}
	switch server.Util.Stage[Message.From.ID] {
	case server.Addnews:
		return t.Execute("/addnews", Message)
	case server.Addquestion:
		return t.Execute("/addquestion", Message)
	case server.Sendanswerto: //Ответ на вопрос
		return t.Execute("/sendanswerto", Message)
	case server.Addmanual:
		return t.Execute("/addmanual", Message)
	default:
		log.Println("default message")
		return t.Routing(Message)

	}
}

func (t *TelegramUpdater) Routing(Message types.TelegramMessage) ([]byte, error) {
	//Определяем есть ли в сообщении сущность команды
	for _, ent := range Message.Entities {
		switch ent.Type {
		case "bot_command":
			log.Println("bot_command")
			return t.Execute(Message.Text, Message)
		default:
			continue
		}
	}
	//
	if Message.Document != nil {
		log.Println("Has document")
	}
	log.Println("no entity")
	return nil, nil
}
