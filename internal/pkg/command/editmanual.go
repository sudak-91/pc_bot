package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type EditManual struct {
	Manuals pubrep.Manuals
	Err     error
}

func (e *EditManual) Handl(Data interface{}) ([]byte, error) {
	query, ok := Data.(types.TelegramCallbackQuery)
	if !ok {
		return nil, fmt.Errorf("EditManual Handl doesn't have TelegramCallbackQuery type in input parametr")
	}
	var Answer types.TelegramSendMessage
	Answer.ChatID = query.From.ID
	Answer.Method = "sendMessage"
	param := strings.Split(query.Data, " ")
	if len(param) != 2 {
		e.Err = errors.New("EditManual Handl has invalid input parametr")
		return util.CommandErrorHandler(&Answer, e.Err)
	}
	Manual, err := e.Manuals.GetManualByID(param[1])
	if err != nil {
		return util.CommandErrorHandler(&Answer, err)
	}
	server.Util.EditManualMutex.Lock()
	server.Util.EditManual[query.From.ID] = Manual
	server.Util.EditManualMutex.Unlock()
	server.Util.StageMutex.Lock()
	server.Util.Stage[query.From.ID] = server.EditManual
	server.Util.StageMutex.Unlock()
	Answer.Text = "Введите новое имя устройства"
	return json.Marshal(Answer)

}
