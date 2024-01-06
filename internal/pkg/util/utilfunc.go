package util

import (
	"encoding/json"
	"log"
	"runtime"

	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

func CommandErrorHandler(Answer *types.SendMessage, err error) ([]byte, error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s : %d has error %s", file, line, err.Error())
	return json.Marshal(*Answer)

}

func CreateAnswer(ID int64) types.SendMessage {
	var Answer types.SendMessage
	Answer.ChatID = ID
	Answer.Method = "sendMessage"
	return Answer
}
