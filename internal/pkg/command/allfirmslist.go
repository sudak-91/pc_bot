package command

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	keyboardmaker "github.com/sudak-91/telegrambotgo/Keyboardmaker"
	methods "github.com/sudak-91/telegrambotgo/TelegramAPI/Methods"
	tgtypes "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type AllFirmsList struct {
	Firms pubrep.Firms
}

func (a *AllFirmsList) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(tgtypes.TelegramCallbackQuery)
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
	Lists, err := a.Firms.GetApprovedFirmsWithOffsetAndLimit(offset, 10)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	keyboard := a.createKeyboard(&Lists, offset)
	var EditMessage methods.EditMessageText
	EditMessage.ChatID = fmt.Sprintf("%s", query.From.ID)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	EditMessage.MessageID = int(query.Message.MessageID)
	EditMessage.ReplyMarkup = keyboard
	EditMessage.Text = "Выберите фирму, чтобы получить список устройств, на которые имеются руководства"
	err = methods.EditMessageTextMethod(EditMessage, os.Getenv("BOT_KEY"))
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	return nil, nil

}
func (a *AllFirmsList) createKeyboard(lists *[]pubrep.Firm, offset int64) *tgtypes.TelegramInlineKeyboardMarkup {
	var keyboard keyboardmaker.InlineCommandKeyboard
	keyboard.MakeGrid(11, 2)
	for k, v := range *lists {
		FirmName := fmt.Sprintf("%s", v.Firm)
		FirmID := v.ID.Hex()
		Command := fmt.Sprintf("/showmanual %s", FirmID)
		keyboard.AddButton(FirmName, Command, k, 0)
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
