package command

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type MarkAsAnswer struct {
	Question repository.Questions
}

func (m *MarkAsAnswer) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("Mark as answer method don't have Callbackquery data type on a input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.ChatID = msg.From.ID
	Answer.Method = "sendMessage"
	if msg.From.ID != server.Util.AdminID {
		Answer.Text = "Не хватает прав"
		log.Println("Markasanswer access currept")
		return json.Marshal(Answer)
	}
	Args := strings.Split(msg.Data, " ")
	if len(Args) != 2 {
		Answer.Text = "Внутреняя ошибка"
		log.Println("Mark as answer need 2 args on data")
		return json.Marshal(Answer)
	}
	var uid uuid.UUID
	id, err := base64.RawURLEncoding.DecodeString(Args[1])
	if err != nil {
		log.Printf("MarkAsAnswer decode uuid has error: %s", err.Error())
		Answer.Text = "Внутренняя ошибка"
		return json.Marshal(Answer)
	}
	for k, v := range id {
		uid[k] = v
	}
	err = m.Question.MarkAsAnswer(uid)
	if err != nil {
		log.Printf("MarkAsAnser handl has repository error: %s\n", err.Error())
		Answer.Text = "Внутреняя ошибка"
		return json.Marshal(Answer)
	}
	Answer.Text = "Вопрос отмечен как прочитанный"
	return json.Marshal(Answer)

}
