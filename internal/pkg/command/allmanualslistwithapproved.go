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
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type AllManualsListWithApproved struct {
	Manuals pubrep.Manuals
}

func (a *AllManualsListWithApproved) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		e := errors.New("Invalid input parametr")
		return nil, e
	}
	answer := util.CreateAnswer(query.From.ID)
	param := strings.Split(query.Data, " ")
	if len(param) != 5 {
		lenerr := errors.New("Invalid input length parametr")
		return util.CommandErrorHandler(&answer, lenerr)
	}
	offset, err := strconv.ParseInt(param[1], 10, 64)
	if err != nil {
		offseterr := errors.New(fmt.Sprintf("Offset parametr can't converted: %w", err))
		return util.CommandErrorHandler(&answer, offseterr)

	}
	approved, err := strconv.ParseBool(param[2])
	if err != nil {
		approvederr := errors.New(fmt.Sprintf("Approved parametr can't converted: %w", err))
		return util.CommandErrorHandler(&answer, approvederr)

	}
	admin, err := strconv.ParseInt(param[3], 10, 64)
	if err != nil {
		adminerr := errors.New(fmt.Sprintf("AdminRole parametr can't converted: %w", err))
		return util.CommandErrorHandler(&answer, adminerr)

	}
	answer.Text = "Выберите мануал"
	UnFirms, err := a.Manuals.GetApprovedManualsWithOffsetAndLimit(param[4], offset, 10, approved)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	if admin == 9 {
		grid := a.createAdminKeyboard(&UnFirms, offset, approved)
		answer.ReplyMarkup = &grid
	} else {
		grid := a.createUserKeyboard(&UnFirms, offset, approved, param[4])
		answer.ReplyMarkup = &grid
	}

	return json.Marshal(answer)

}

func (a *AllManualsListWithApproved) createAdminKeyboard(ManualList *[]pubrep.Manual, offset int64, approved bool) tgtypes.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 3)
	for k, v := range *ManualList {
		if !approved {
			keyboard.AddButton(fmt.Sprintf("%s: Утвердить", v.DeviceModel), fmt.Sprintf("/approvedmanual %s", v.ManualID.Hex()), k, 0)
		}
		keyboard.AddButton(fmt.Sprintf("%s: Редактировать", v.DeviceModel), fmt.Sprintf("/editmanual %s", v.ManualID.Hex()), k, 1)
		keyboard.AddButton(fmt.Sprintf("%s: Удалить", v.DeviceModel), fmt.Sprintf("/deletemanual %s", v.ManualID.Hex()), k, 2)
	}
	if offset != 0 {
		Backcommand := fmt.Sprintf("/allmanualslistwithapproved %d %t %d", offset-10, approved, 9)
		keyboard.AddButton("Назад", Backcommand, 10, 0)
	}
	if len(*ManualList) == 10 {
		ForwardCommand := fmt.Sprintf("/allmanualslistwithapproved %d %t %d", offset+10, approved, 9)
		keyboard.AddButton("Вперед", ForwardCommand, 10, 2)
	}
	grid := keyboard.GetKeyboard()
	return grid
}

func (a *AllManualsListWithApproved) createUserKeyboard(ManualList *[]pubrep.Manual, offset int64, approved bool, firmid string) tgtypes.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 2)
	for k, v := range *ManualList {
		keyboard.AddButton(fmt.Sprintf("Показать руководство%s", v.DeviceModel), fmt.Sprintf("/showmanual %s", v.ManualID.Hex()), k, 0)

	}
	if offset != 0 {
		Backcommand := fmt.Sprintf("/allmanualslistwithapproved %d %t %d %s", offset-10, approved, 0, firmid)
		keyboard.AddButton("Назад", Backcommand, 10, 0)
	}
	if len(*ManualList) == 10 {
		ForwardCommand := fmt.Sprintf("/allmanualslistwithapproved %d %t %d %s", offset+10, approved, 0, firmid)
		keyboard.AddButton("Вперед", ForwardCommand, 10, 1)
	}
	grid := keyboard.GetKeyboard()
	return grid
}
