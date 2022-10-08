package command

import (
	"encoding/json"
	"errors"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type ManualArchive struct {
}

func (m *ManualArchive) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, errors.New("Invalid input parametr")
	}
	answer := util.CreateAnswer(query.From.ID)
	answer.ReplyMarkup = m.createKeyboard()
	return json.Marshal(answer)

}

func (m *ManualArchive) createKeyboard() *types.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(3, 1)
	keyboard.AddButton("Показать список всех фирм", "/allfirmlist", 0, 0) //TODO:Реализовать метод
	keyboard.AddButton("Алфавитный указатель", "/alphabetlist", 0, 1)     //TODO:
	keyboard.AddButton("Поиск", "/search", 0, 2)                          //TODO:
	rslt := keyboard.GetKeyboard()
	return &rslt

}
