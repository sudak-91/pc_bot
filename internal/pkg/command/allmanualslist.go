package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	"github.com/sudak-91/telegrambotgo/pkg/telegramerrors"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type AllManualsList struct {
	Manuals pubrep.Manuals
}

func (a *AllManualsList) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, &telegramerrors.InvalidInputParametrType{}
	}
	answer := util.CreateAnswer(query.From.ID)
	param := strings.Split(query.Data, " ")
	if len(param) != 2 {
		lenerror := telegramerrors.NewInvalidInputParametrLength(2, int32(len(param)))
		return nil, lenerror
	}
	FirmId := param[1]
	Result, err := a.Manuals.GetManualsByFirmID(FirmId)
	if err != nil {
		return nil, err
	}
	answer.Text = "Выберите действие с мануалом"
	answer.ReplyMarkup = a.createKeyboard(&Result)
	return json.Marshal(answer)

}

func (a *AllManualsList) createKeyboard(Manuals *[]pubrep.Manual) *types.InlineKeyboardMarkup {
	keyboard := keyboardmaker.InlineCommandKeyboard{}
	keyboard.MakeGrid(len(*Manuals), 3)
	for k, v := range *Manuals {
		keyboard.AddButton(v.DeviceModel, fmt.Sprint("/showmanual %s", v.ManualID.Hex()), k, 0)
		keyboard.AddButton(fmt.Sprint("Удалить %s", v.DeviceModel), fmt.Sprintf("/deletemanual %s", v.ManualID.Hex()), k, 1)
		if !v.Approved {
			buttontext := fmt.Sprintf("Подтвердить %s", v.DeviceModel)
			command := fmt.Sprintf("/approvedmanual %s", v.ManualID.Hex())
			keyboard.AddButton(buttontext, command, k, 2)
		}
	}
	markup := keyboard.GetKeyboard()
	return &markup
}
