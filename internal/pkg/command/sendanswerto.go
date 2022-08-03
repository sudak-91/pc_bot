package command

//TODO:SendAnswerTo отправлят ответ конкретному контрибутеру
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/pkg/repository"
	methods "github.com/sudak-91/telegrambotgo/TelegramAPI/Methods"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type SendAnswerTo struct {
	Question repository.Questions
}

func (s *SendAnswerTo) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("SendAnswerTo dont has Telegram message on the input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.ChatID = msg.From.ID
	Answer.Method = "sendMessage"

	var sMessage methods.SendMessage
	sMessage.ChatID = server.Util.AnswerCtx[msg.From.ID].ContributerID
	sMessage.ReplyToMessageID = server.Util.AnswerCtx[msg.From.ID].MessageID
	sMessage.Text = msg.Text
	err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), sMessage)
	if err != nil {
		log.Printf("SendAnswerTo has SendMessageMethod error: %s", err.Error())
		delete(server.Util.AnswerCtx, msg.From.ID)
		delete(server.Util.Stage, msg.From.ID)
		Answer.Text = "Внутренняя ошибка"
		return json.Marshal(Answer)
	}

	id, err := base64.RawStdEncoding.DecodeString(server.Util.AnswerCtx[msg.From.ID].QuestionID)
	if err != nil {
		Answer.Text = "Внутреняя ошибка"
		delete(server.Util.AnswerCtx, msg.From.ID)
		delete(server.Util.Stage, msg.From.ID)
		return json.Marshal(Answer)
	}
	//FIXME: Переделать  систему с UUID
	var uid uuid.UUID
	for k, v := range id {
		uid[k] = v
	}
	err = s.Question.MarkAsAnswer(uid)
	if err != nil {
		Answer.Text = "Внутреняя ошибка"
		log.Printf("SendAnswerTo has SendMessageMethod error: %s", err.Error())
		delete(server.Util.AnswerCtx, msg.From.ID)
		delete(server.Util.Stage, msg.From.ID)
		return json.Marshal(Answer)
	}
	Answer.Text = "Ответ отправлен"
	delete(server.Util.AnswerCtx, msg.From.ID)
	delete(server.Util.Stage, msg.From.ID)

	return json.Marshal(Answer)
}
