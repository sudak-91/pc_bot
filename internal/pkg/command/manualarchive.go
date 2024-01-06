package command

import (
	"encoding/json"
	"errors"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type ManualArchive struct {
}

func (m *ManualArchive) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, errors.New("Invalid input parametr")
	}
	answer := util.CreateAnswer(query.From.ID)
	answer.ReplyMarkup = m.createKeyboard()
	return json.Marshal(answer)

}

func (m *ManualArchive) createKeyboard() *types.InlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(3, 1)
	keyboard.AddButton("Показать список всех фирм", "/allfirmslist 0", 0, 0) //TODO:Реализовать метод
	keyboard.AddButton("Алфавитный указатель", "/alphabetlist", 1, 0)        //TODO:
	keyboard.AddButton("Поиск", "/search", 2, 0)                             //TODO:
	rslt := keyboard.GetKeyboard()
	return &rslt

}
