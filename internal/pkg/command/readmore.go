package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type ReadMore struct {
	News pubrep.Newser
}

func (r *ReadMore) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("The Readmore handl dont have Callbackquery type on a input parametr\n")

	}
	Args := strings.Split(msg.Data, " ")
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	if len(Args) != 2 {
		log.Println("Internal Error")
		Answer.Text = "Произошла внутреняя ошибка"
		return json.Marshal(Answer)
	}
	News, err := r.News.GetNews(Args[1])
	if err != nil {
		log.Printf("ReadMore Handl has error:%s/n", err.Error())
		Answer.Text = "Внутреняя ошибка"
		return json.Marshal(Answer)
	}
	Answer.Text = News[0].Text
	var newsKeyboard keyboardmaker.InlineCommandKeyboard
	newsKeyboard.MakeGrid(1, 1)
	newsKeyboard.AddButton("Отметить как прочитанное", fmt.Sprintf("/markasread %v", Args[1]), 0, 0)
	kboard := newsKeyboard.GetKeyboard()
	Answer.ReplyMarkup = &kboard
	return json.Marshal(Answer)
}
