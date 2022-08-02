package command

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type Login struct {
	Users    pubrep.Users
	Keyboard types.TelegramInlineKeyboardMarkup
}

func (l *Login) Handl(data interface{}) ([]byte, error) {
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("Login command don have message type on enter parametr")
	}
	Answer.ChatID = msg.Chat.ID
	Args := strings.Split(msg.Text, " ")
	if len(Args) != 2 {
		Answer.Text = "Неверное количество параметров"
		return json.Marshal(Answer)
	}
	if Args[1] != os.Getenv("ADMIN_PASS") {
		Answer.Text = "Неверный пароль"
		return json.Marshal(Answer)
	}
	usr, err := l.Users.GetUser(msg.From.ID)
	if err != nil {
		return nil, err
	}
	usr[0].Role = 9
	err = l.Users.UpdateUser(usr[0])
	if err != nil {
		return nil, err
	}
	Answer.Text = "Вы получили права администратора"
	Answer.ReplyMarkup = &l.Keyboard
	server.Util.AdminID = msg.Chat.ID
	return json.Marshal(Answer)
}
