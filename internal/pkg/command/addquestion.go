package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	methods "github.com/sudak-91/telegrambotgo/TelegramAPI/Methods"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type AddQuestion struct {
	Question pubrep.Questions
}

func (q *AddQuestion) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("AddQuestion handl dont have Telegram message type on the input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.Chat.ID
	Answer.Text = "Ваш вопрос добавлен"
	if err := q.Question.CreateQuestion(msg.Text, msg.Chat.ID, int64(msg.MessageID)); err != nil {
		Answer.Text = "Произошла внутренняя ошибка"
		log.Println(err.Error())
	}
	var sMessage methods.SendMessage
	sMessage.ChatID = server.Util.AdminID
	sMessage.Text = "Добавлен новый вопрос"
	err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), sMessage)
	if err != nil {
		log.Printf("Send message has error: %s", err)
	}
	delete(server.Util.Stage, msg.Chat.ID)
	return json.Marshal(Answer)
}
