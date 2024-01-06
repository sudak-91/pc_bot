package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type ApprovedFirm struct {
	Firms   pubrep.Firms
	Manuals pubrep.Manuals
}

func (a *ApprovedFirm) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, fmt.Errorf("Internal error")
	}
	var Answer types.SendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = query.From.ID
	param := strings.Split(query.Data, " ")
	if len(param) != 2 {
		return util.CommandErrorHandler(&Answer, fmt.Errorf("Неверное количество параметров"))
	}
	manuals, err := a.Firms.GetFirmById(param[1])
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	manuals.Approved = true
	err = a.Firms.UpdateFirm(manuals)
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	err = a.Manuals.UpdateEmbeddedFirm(manuals)
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	Answer.Text = "Фирма подтверждена"
	return json.Marshal(Answer)
}
