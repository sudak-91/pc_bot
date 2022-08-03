package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type SendAnswer struct {
	Questions repository.Questions
}

// data: QuestionID, MessageID, ContributerID
func (s *SendAnswer) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("SendAnswer handle dont has Callbackquery on a iput parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	if msg.From.ID != server.Util.AdminID {
		Answer.Text = "У вас недостаточно прав"
		log.Println("Access currepted")
		return json.Marshal(Answer)
	}
	Args := strings.Split(msg.Data, " ")
	if len(Args) != 4 {
		log.Println("Data dont have 4 parametr")
		Answer.Text = "Внутренняя ошибка"
		return json.Marshal(Answer)
	}

	MessageID, err := strconv.Atoi(Args[2])
	if err != nil {
		log.Println("Convert error: %s", err.Error())
		Answer.Text = "Внутреняя ошибка"
		return json.Marshal(Answer)
	}
	ContributerID, err := strconv.ParseInt(Args[3], 10, 64)
	if err != nil {
		log.Println("Convert error: %s", err.Error())
		Answer.Text = "Внутреняя ошибка"
		return json.Marshal(Answer)
	}

	var Ctx server.SendAnswer
	Ctx.ContributerID = ContributerID
	Ctx.MessageID = int32(MessageID)
	Ctx.QuestionID = Args[1]
	server.Util.AnswerCtx[msg.From.ID] = Ctx
	server.Util.Stage[msg.From.ID] = 30

	Answer.Text = "Далее ваш ответ будет отправлен тому кто задал вопрос"
	return json.Marshal(Answer)
}
