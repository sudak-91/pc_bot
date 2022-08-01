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
	err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), server.Util.AdminID, "добавлен новый вопрос")
	if err != nil {
		log.Println("Send message has error: %s", err)
	}
	return json.Marshal(Answer)
}
