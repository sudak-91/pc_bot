package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/server"
	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type ConfirmEditFirm struct {
	Firms pubrep.Firms
}

func (c *ConfirmEditFirm) Handl(data interface{}) ([]byte, error) {
	msg, ok := data.(types.TelegramMessage)
	if !ok {
		return nil, fmt.Errorf("Internal error")
	}
	var Answer types.TelegramSendMessage
	Answer.Method = "sendMessage"
	Answer.ChatID = msg.From.ID
	NewName := strings.Split(msg.Text, " ")
	if len(NewName) != 1 {
		delete(server.Util.Stage, msg.From.ID)
		delete(server.Util.EditFirm, msg.Chat.ID)
		return util.CommandErrorHandler(&Answer, fmt.Errorf("Internal error"))
	}
	objFirm, ok := server.Util.EditFirm[msg.Chat.ID]
	if !ok {
		delete(server.Util.Stage, msg.From.ID)
		delete(server.Util.EditFirm, msg.Chat.ID)
		return util.CommandErrorHandler(&Answer, fmt.Errorf("Internal error"))
	}
	delete(server.Util.EditFirm, msg.Chat.ID)
	objFirm.Firm = NewName[0]
	err := c.Firms.UpdateFirm(objFirm)
	if err != nil {
		delete(server.Util.Stage, msg.From.ID)
		delete(server.Util.EditFirm, msg.Chat.ID)
		return util.CommandErrorHandler(&Answer, err)
	}
	Answer.Text = "Обновление завершено"
	delete(server.Util.Stage, msg.From.ID)
	delete(server.Util.EditFirm, msg.Chat.ID)
	return json.Marshal(Answer)
}
