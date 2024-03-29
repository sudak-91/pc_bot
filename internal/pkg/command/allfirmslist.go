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
	tgtypes "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type AllFirmsList struct {
	Firms pubrep.Firms
}

func (a *AllFirmsList) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(tgtypes.CallbackQuery)
	log.Printf("%+v", query)
	if !ok {
		return nil, errors.New("Invalid input parametr")
	}
	param := strings.Split(query.Data, " ")
	answer := util.CreateAnswer(query.From.ID)
	if len(param) != 2 {
		lenerr := errors.New("Invalid data length")

		return util.CommandErrorHandler(&answer, lenerr)
	}
	offset, err := strconv.ParseInt(param[1], 10, 32)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	Lists, err := a.Firms.GetAllFirmsWithOffsetAndLimit(offset, 10)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	answer.Text = "Выберите фирму"
	keyboard := a.createKeyboard(&Lists, offset)
	//var EditMessage methods.EditMessageText
	/*EditMessage.ChatID = fmt.Sprintf("%s", query.From.ID)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	EditMessage.MessageID = int(query.Message.From.ID)
	EditMessage.ReplyMarkup = keyboard
	EditMessage.Text = "Выберите фирму, чтобы получить список устройств, на которые имеются руководства"
	err = methods.EditMessageTextMethod(EditMessage, os.Getenv("BOT_KEY"))
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}*/
	answer.ReplyMarkup = keyboard
	return json.Marshal(answer)

}
func (a *AllFirmsList) createKeyboard(lists *[]pubrep.Firm, offset int64) *tgtypes.InlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 2)
	for k, v := range *lists {
		FirmName := fmt.Sprintf("%s", v.Firm)
		FirmID := v.ID.Hex()
		Command := fmt.Sprintf("/allmanualslist %s", FirmID)
		keyboard.AddButton(FirmName, Command, k, 0)
		DeleteCommand := fmt.Sprintf("/deletefirm %s", FirmID)
		keyboard.AddButton(fmt.Sprintf("Удалить: %s", FirmName), DeleteCommand, k, 1)
	}
	if offset != 0 {
		BackCommand := fmt.Sprintf("/allfirmslist %d", offset-10)
		keyboard.AddButton("Назад", BackCommand, 10, 0)
	}
	if len(*lists) == 10 {
		ForwardCommand := fmt.Sprintf("/allfirmslist %d", offset+10)
		keyboard.AddButton("Вперед", ForwardCommand, 10, 1)
	}
	grid := keyboard.GetKeyboard()
	return &grid

}
