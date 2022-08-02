package command

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	methods "github.com/sudak-91/telegrambotgo/TelegramAPI/Methods"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type Shown struct {
	News pubrep.Newser
}

func (s *Shown) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("shown handl dont have CallbackQuery data type on the input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	if msg.From.ID != server.Util.AdminID {
		Answer.Text = "У вас недостаточно прав для использования этой команды"
		return json.Marshal(Answer)
	}
	News, err := s.News.GetNotAsReadNews()
	if err != nil {
		log.Printf("Shown handle has error :%s", err.Error())
		Answer.Text = "Произошла внутреняя ошибка"
		return json.Marshal(Answer)
	}
	for k, v := range News {
		var sMessage methods.SendMessage
		sMessage.ChatID = msg.From.ID
		sMessage.Text = v.Text[:100]

	}
}
