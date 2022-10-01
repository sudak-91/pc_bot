package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type ApprovedManual struct {
	Manual pubrep.Manuals
}

func (a *ApprovedManual) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("ApprovedManual has invalid input data")
	}
	Answer := util.CreateAnswer(query.From.ID)
	param := strings.Split(query.Data, " ")
	if len(param) != 2 {
		return util.CommandErrorHandler(&Answer, fmt.Errorf("Invalid input parametr count"))
	}
	Manual, err := a.Manual.GetManualByID(param[1])
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	Manual.Approved = true
	err = a.Manual.UpdateManual(Manual)
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	Answer.Text = "Руководство подвтерждено"
	return json.Marshal(Answer)
}
