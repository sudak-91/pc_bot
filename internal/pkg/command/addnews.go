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

type AddNews struct {
	News pubrep.Newser
}

func (a *AddNews) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("AddNews handler dont have TelegramMessage type on input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.Chat.ID
	if err := a.News.CreateNews(msg.Text, msg.Chat.ID); err != nil {
		Answer.Text = "Произошел сбой при добавлении новости."
		log.Println(err.Error())
		return json.Marshal(Answer)
	}
	methods.SendMessageMethod(os.Getenv("BOT_KEY"), server.Util.AdminID, "Добавлена новая новость")
	return json.Marshal(Answer)
}
