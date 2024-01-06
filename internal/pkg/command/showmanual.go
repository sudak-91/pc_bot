package command

import (
	"os"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"github.com/sudak-91/telegrambotgo/pkg/telegramerrors"
	methods "github.com/sudak-91/telegrambotgo/telegram_api/methods"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type ShowManual struct {
	Manuals pubrep.Manuals
}

func (s *ShowManual) Handl(data interface{}) ([]byte, error) {
	qry, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, &telegramerrors.InvalidInputParametrType{}
	}
	answer := util.CreateAnswer(qry.From.ID)
	param := strings.Split(qry.Data, " ")
	if len(param) != 2 {
		return util.CommandErrorHandler(&answer, telegramerrors.NewInvalidInputParametrLength(2, int32(len(param))))
	}
	ManualID := param[1]
	rslt, err := s.Manuals.GetManualByID(ManualID)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	sendDoc := methods.SendDocument{}
	sendDoc.ChatId = qry.From.ID
	sendDoc.Document = rslt.FileUniqID
	err = methods.SendDocumentMethod(os.Getenv("BOT_KEY"), sendDoc)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	return nil, nil

}
