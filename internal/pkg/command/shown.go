package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
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
	for _, v := range News {
		var sMessage methods.SendMessage
		sMessage.ChatID = msg.From.ID
		sMessage.Text = v.Text
		if len(v.Text) > 140 {
			sMessage.Text = v.Text[:140]
		}
		var newsKeyboard keyboardmaker.InlineCommandKeyboard
		newsKeyboard.MakeGrid(1, 2)
		//FIXME: Исправить выдачу кнопок
		log.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@")
		log.Println(v.NewsID)
		q := fmt.Sprintf("/readmore %v", v.NewsID[:])
		log.Println(q)
		newsKeyboard.AddButton("Прочесть полностью", fmt.Sprintf("/readmore %s", v.NewsID[:]), 0, 0)
		newsKeyboard.AddButton("Отметить как прочитанное", fmt.Sprintf("/markasread %s", v.NewsID[:]), 0, 1)
		kboard := newsKeyboard.GetKeyboard()
		sMessage.ReplayMarkup = &kboard
		if err := methods.SendMessageMethod(os.Getenv("BOT_KEY"), sMessage); err != nil {
			log.Printf("send message on shown handl has error:%s\n", err.Error())
			continue
		}
	}
	Answer.Text = "Выдача новостей завершена"
	return json.Marshal(Answer)
}
