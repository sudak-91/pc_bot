package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type EditFirmCommand struct {
	Firms pubrep.Firms
}

func (e *EditFirmCommand) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("Input parametr is not TelegramCallnackquery type")
	}
	log.Printf("%s\n", query.Data)
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = query.From.ID

	param := strings.Split(query.Data, " ")
	log.Printf("%v\n", param)
	if len(param) > 0 {
		return util.CommandErrorHandler(&Answer, fmt.Errorf("Input Parametr Error"))
	}
	FirmData, err := e.Firms.GetFirmById(query.Data)
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	server.Util.EditFirm[query.From.ID] = FirmData[0]
	Answer.Text = fmt.Sprintf("Cтарое название: %s\n Введите исправленное название", FirmData[0].Firm)
	server.Util.Stage[query.From.ID] = server.EditFirm
	return json.Marshal(Answer)

}
