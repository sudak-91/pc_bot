package command

import (
	"encoding/json"
	"fmt"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

//Добавление новой стадии для пользователя в кэщ
type News struct {
}

func (n *News) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("News command dont have TelegramCallbackQuery on input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	server.Util.Stage[msg.From.ID] = 10
	Answer.Text = `Далее можете предложить свою новость. Не забудьте оствить ссылку на первоисточник и краткое описание. Укажите ваш социальные сети, если хотите, чтобы мы их опубликовали`
	return json.Marshal(Answer)
}
