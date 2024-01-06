package command

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	methods "github.com/sudak-91/telegrambotgo/telegram_api/methods"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type ShowQ struct {
	Question pubrep.Questions
}

func (s *ShowQ) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, fmt.Errorf("ShowQ Handl dont have CallbackQuery data type on a input parametr")
	}
	var Answer types.SendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	if msg.From.ID != server.Util.AdminID {
		Answer.Text = "У вас недостаточно прав"
		log.Println("Access currapted")
		return json.Marshal(Answer)
	}
	Questions, err := s.Question.GetNotAnswerQuestion()
	if err != nil {
		log.Printf("ShowQ handl has repository error: %s", err.Error())
		Answer.Text = "Произошла внутренняя ошибка"
		return json.Marshal(Answer)
	}
	for _, v := range Questions {
		var sMessage methods.SendMessage
		sMessage.Text = v.Text
		var qkeyboar keyboardmaker.InlineCommandKeyboard
		qkeyboar.MakeGrid(1, 2)
		an := fmt.Sprintf("/sendanswer %s %d %d", base64.RawURLEncoding.EncodeToString(v.QuestionID[:]), v.MessageID, v.ContributerID)
		qkeyboar.AddButton("Отправить ответ", an, 0, 0)
		mkaa := fmt.Sprintf("/markasanswer %s", base64.RawURLEncoding.EncodeToString(v.QuestionID[:]))
		qkeyboar.AddButton("Отметить как прочитаное", mkaa, 0, 1)
		kb := qkeyboar.GetKeyboard()
		sMessage.ReplayMarkup = &kb
		sMessage.ChatID = msg.From.ID
		err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), sMessage)
		if err != nil {
			log.Printf("ShowQ Handl has send message error: %s", err.Error())
			continue
		}
	}
	Answer.Text = "Все вопросы получены"
	return json.Marshal(Answer)
}
