package service

import (
	"log"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

// TODO: Обработка ошибок
func (t *TelegramUpdater) messageService(Message types.Message) ([]byte, error) {
	_, ok := server.Util.Stage[Message.From.ID]
	if !ok {
		return t.Routing(Message)
	}
	switch server.Util.Stage[Message.From.ID] {
	case server.Addnews:
		return t.Execute("/addnews", Message)
	case server.AddQuestion:
		return t.Execute("/addquestion", Message)
	case server.SendAnswerTo: //Ответ на вопрос
		return t.Execute("/sendanswerto", Message)
	case server.AddManual:
		return t.Execute("/addmanual", Message)
	case server.AddManualInfo:
		return t.Execute("/addmanualinfo", Message)
	case server.AddManualDocument:
		return t.Execute("/addmanualdocument", Message)
	case server.EditFirm:
		return t.Execute("/confirmeditfirm", Message)
	case server.EditManual:
		return t.Execute("/confirmeditmanual", Message)
	default:
		log.Println("default message")
		return t.Routing(Message)

	}
}

func (t *TelegramUpdater) Routing(Message types.Message) ([]byte, error) {
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
