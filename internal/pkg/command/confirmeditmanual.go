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

type ConfirmEditManual struct {
	Manual pubrep.Manuals
	Err    error
}

func (c *ConfirmEditManual) Handl(data interface{}) ([]byte, error) {
	Message, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("ConfirmEditManual Handl has invalid input parametr")
	}
	Answer := util.CreateAnswer(Message.From.ID)
	param := strings.Split(Message.Text, " ")
	server.Util.EditManualMutex.Lock()
	server.Util.StageMutex.Lock()
	defer func() {
		server.Util.StageMutex.Unlock()
		server.Util.EditManualMutex.Unlock()
	}()
	if len(param) != 1 {
		c.Err = errors.New("ConfirmEditManaul Handl has invalid input parametr")
		delete(server.Util.EditManual, Message.From.ID)
		delete(server.Util.Stage, Message.From.ID)
		return util.CommandErrorHandler(&Answer, c.Err)
	}

	Manual := server.Util.EditManual[Message.From.ID]
	Manual.DeviceModel = param[0]
	err := c.Manual.UpdateManual(Manual)
	if err != nil {
		delete(server.Util.EditManual, Message.From.ID)
		delete(server.Util.Stage, Message.From.ID)
		return util.CommandErrorHandler(&Answer, err)
	}
	Answer.Text = "Обновление завершено"
	delete(server.Util.EditManual, Message.From.ID)
	delete(server.Util.Stage, Message.From.ID)
	return json.Marshal(Answer)
}
