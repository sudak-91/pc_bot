package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type EditFirmCommand struct {
	Firms pubrep.Firms
}

func (e *EditFirmCommand) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, fmt.Errorf("Input parametr is not TelegramCallnackquery type")
	}
	log.Printf("%s\n", query.Data)
	var Answer types.SendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = query.From.ID

	param := strings.Split(query.Data, " ")
	log.Printf("%v\n", param)
	if len(param) != 2 {
		return util.CommandErrorHandler(&Answer, fmt.Errorf("Input Parametr Error"))
	}
	FirmData, err := e.Firms.GetFirmById(param[1])
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	server.Util.EditFirm[query.From.ID] = FirmData
	Answer.Text = fmt.Sprintf("Cтарое название: %s\n Введите исправленное название", FirmData.Firm)
	server.Util.Stage[query.From.ID] = server.EditFirm
	return json.Marshal(Answer)

}
