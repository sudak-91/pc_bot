package command

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/sudak-91/pc_bot/internal/pkg/util"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	tgerros "github.com/sudak-91/telegrambotgo/pkg/telegramerrors"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteFirm struct {
	Firms  pubrep.Firms
	Manual pubrep.Manuals
	Client *mongo.Client
}

func (d *DeleteFirm) Handl(data interface{}) ([]byte, error) {
	query, ok := data.(types.CallbackQuery)
	if !ok {
		return nil, &tgerros.InvalidInputParametrType{}
	}
	answer := util.CreateAnswer(query.From.ID)
	paramters := strings.Split(query.Data, " ")
	if len(paramters) != 2 {
		answerlen := len(paramters)
		return util.CommandErrorHandler(&answer, tgerros.NewInvalidInputParametrLength(2, int32(answerlen)))
	}
	FirmID := paramters[1]

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		if err := d.Firms.DeleteFirm(FirmID); err != nil {
			return nil, err
		}
		if err := d.Manual.DeleteManualsByFirm(FirmID); err != nil {
			return nil, err
		}
		return nil, nil
	}
	session, err := d.Client.StartSession()
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	defer session.EndSession(context.TODO())
	result, err := session.WithTransaction(context.TODO(), callback)
	if err != nil {
		return util.CommandErrorHandler(&answer, err)
	}
	log.Println(result)
	answer.Text = "Фирма и все мануалы удалены"
	return json.Marshal(answer)

}
