package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	tgtypes "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type AllUnapprovedFirmsList struct {
	Firms pubrep.Firms
}

func (a *AllUnapprovedFirmsList) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(tgtypes.TelegramCallbackQuery)
	if !ok {
		e := errors.New("Invalid input parametr")
		return nil, e
	}
	answer := util.CreateAnswer(query.From.ID)
	param := strings.Split(query.Data, " ")
	if len(param) != 2 {
		lenerr := errors.New("Invalid input length parametr")
		return util.CommandErrorHandler(&answer, lenerr)
	}
	offset, err := strconv.ParseInt(param[1], 10, 64)
	if err != nil {
		offseterr := errors.New(fmt.Sprintf("Offset parametr can't converted: %w", err))
		return util.CommandErrorHandler(&answer, offseterr)

	}
	UnFirms, err := a.Firms.GetApprovedFirmsWithOffsetAndLimit(offset, 10, false)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	grid := a.createKeyboard(&UnFirms, offset)
	answer.ReplyMarkup = &grid
	return json.Marshal(answer)

}

func (a *AllUnapprovedFirmsList) createKeyboard(FirmsList *[]pubrep.Firm, offset int64) tgtypes.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 3)
	for k, v := range *FirmsList {
		keyboard.AddButton(fmt.Sprintf("%s: Утвердить"), fmt.Sprintf("/approvedfirm %s", v.ID.Hex()), k, 0)
		keyboard.AddButton(fmt.Sprintf("%s: Редактировать"), fmt.Sprintf("/editfirm %s", v.ID.Hex()), k, 1)
		keyboard.AddButton(fmt.Sprintf("%s: Удалить"), fmt.Sprintf("/deletefirm %s", v.ID.Hex()), k, 2)
	}
	if offset != 0 {
		Backcommand := fmt.Sprintf("/allunapprovedfirmslist %d", offset-10)
		keyboard.AddButton("Назад", Backcommand, 10, 0)
	}
	if len(*FirmsList) == 10 {
		ForwardCommand := fmt.Sprintf("/allunapprovedfirmslist %d", offset+10)
		keyboard.AddButton("Вперед", ForwardCommand, 10, 2)
	}
	grid := keyboard.GetKeyboard()
	return grid
}
