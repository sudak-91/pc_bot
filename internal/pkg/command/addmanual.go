package command

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type AddManual struct {
	Manual pubrep.Manuals
}

//TODO: Изменить метод. Требуется для начала выводить всю информацию и разрешить ее редактирование
func (m *AddManual) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("Handl AddManula has error")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	if msg.Document == nil {

		log.Print("Empty document")
		Answer.Text = "Нет документа в сообщении"
		return json.Marshal(Answer)
	}
	v, ok := server.Util.Manual[msg.From.ID]
	if !ok {

		log.Print("Empty document")
		Answer.Text = "Нет необходимых данных. Попробуйте начать заного"
		delete(server.Util.Stage, msg.From.ID)
		return json.Marshal(Answer)
	}
	v.FileUniqID = msg.Document.FileUniqueID
	if err := m.Manual.CreateManual(v); err != nil {
		Answer.Text = "Внутренняя ошибка. Попробуйте снова"

		log.Print("Empty document")
		delete(server.Util.Stage, msg.From.ID)
		delete(server.Util.Manual, msg.From.ID)
		return json.Marshal(Answer)
	}
	Answer.Text = "Руководство добавлено"
	return json.Marshal(Answer)

}
