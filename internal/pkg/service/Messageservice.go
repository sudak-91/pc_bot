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
	case int(addnews):
		return t.Execute("/addnews", Message)
	case int(addquestion):
		return t.Execute("/addquestion", Message)
	case int(sendanswerto): //Ответ на вопрос
		return t.Execute("/sendanswerto", Message)
	case int(addmanual):
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
