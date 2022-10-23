package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	tgtypes "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

//Get firms list with approved flag and has limit and offset
type AllFirmsListWithApproved struct {
	Firms pubrep.Firms
}

func (a *AllFirmsListWithApproved) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(tgtypes.TelegramCallbackQuery)
	if !ok {
		e := errors.New("Invalid input parametr")
		return nil, e
	}
	log.Printf("%v\n", query)
	answer := util.CreateAnswer(query.From.ID)
	param := strings.Split(query.Data, " ")
	if len(param) != 4 {
		lenerr := errors.New("Invalid input length parametr")
		return util.CommandErrorHandler(&answer, lenerr)
	}
	log.Printf("%v\n", param)
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

	UnFirms, err := a.Firms.GetApprovedFirmsWithOffsetAndLimit(offset, 10, approved)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	if admin == 9 {
		grid := a.createAdminKeyboard(&UnFirms, offset, approved)
		answer.ReplyMarkup = &grid
	} else {
		grid := a.createUserKeyboard(&UnFirms, offset, approved)
		answer.ReplyMarkup = &grid
	}

	return json.Marshal(answer)

}

func (a *AllFirmsListWithApproved) createAdminKeyboard(FirmsList *[]pubrep.Firm, offset int64, approved bool) tgtypes.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 3)
	for k, v := range *FirmsList {
		if !approved {
			keyboard.AddButton(fmt.Sprintf("%s: Утвердить", v.Firm), fmt.Sprintf("/approvedfirm %s", v.ID.Hex()), k, 0)
		}
		keyboard.AddButton(fmt.Sprintf("%s: Редактировать", v.Firm), fmt.Sprintf("/editfirm %s", v.ID.Hex()), k, 1)
		keyboard.AddButton(fmt.Sprintf("%s: Удалить", v.Firm), fmt.Sprintf("/deletefirm %s", v.ID.Hex()), k, 2)
	}
	if offset != 0 {
		Backcommand := fmt.Sprintf("/allfirmslistwithapproved %d %t %d", offset-10, approved, 9)
		keyboard.AddButton("Назад", Backcommand, 10, 0)
	}
	if len(*FirmsList) == 10 {
		ForwardCommand := fmt.Sprintf("/allfirmslistwithapproved %d %t %d", offset+10, approved, 9)
		keyboard.AddButton("Вперед", ForwardCommand, 10, 2)
	}
	grid := keyboard.GetKeyboard()
	return grid
}

func (a *AllFirmsListWithApproved) createUserKeyboard(FirmsList *[]pubrep.Firm, offset int64, approved bool) tgtypes.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 2)
	for k, v := range *FirmsList {
		keyboard.AddButton(fmt.Sprintf("Показать список руководст %s", v.Firm), fmt.Sprintf("/allmanualslistwithapproved %s", v.ID.Hex()), k, 0)

	}
	if offset != 0 {
		Backcommand := fmt.Sprintf("/allfirmslistwithapproved %d %t %d", offset-10, approved, 0)
		keyboard.AddButton("Назад", Backcommand, 10, 0)
	}
	if len(*FirmsList) == 10 {
		ForwardCommand := fmt.Sprintf("/allfirmslistwithapproved %d %t %d", offset+10, approved, 0)
		keyboard.AddButton("Вперед", ForwardCommand, 10, 1)
	}
	grid := keyboard.GetKeyboard()
	return grid
}
