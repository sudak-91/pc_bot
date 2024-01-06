package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type MarkAsRead struct {
	News pubrep.Newser
}

func (m *MarkAsRead) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, fmt.Errorf("The MarkAsRead handle dont have Callback Query data type on the input parameter")
	}
	var Answer types.SendMessage
	log.Printf("Message text is: %s\n", msg.Message.Text)
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	Args := strings.Split(msg.Data, " ")
	if len(Args) != 2 {
		Answer.Text = "Произошла внутренняя ошибка"
		log.Println("Arguments count more then 2")
		return json.Marshal(Answer)
	}
	News, err := m.News.GetNews(Args[1])
	if err != nil {
		log.Printf("MarkAsRead handle has error: %s\n", err.Error())
		Answer.Text = "Произоша внутрення ошибка"
		return json.Marshal(Answer)
	}
	News[0].AsRead = true
	err = m.News.UpdateNews(News[0])
	if err != nil {
		log.Printf("MarkAsRead handle has error: %s\n", err.Error())
		Answer.Text = "Произоша внутрення ошибка"
		return json.Marshal(Answer)
	}
	Answer.Text = "Новость успешно обновлена"
	return json.Marshal(Answer)
}
