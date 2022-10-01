package util

import (
	"encoding/json"
	"log"
	"runtime"

	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

func CommandErrorHandler(Answer *types.TelegramSendMessage, err error) ([]byte, error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s : %d has error %s", file, line, err.Error())
	return json.Marshal(*Answer)

}

func CreateAnswer(ID int64) types.TelegramSendMessage {
	var Answer types.TelegramSendMessage
	Answer.ChatID = ID
	Answer.Method = "sendMessage"
	return Answer
}
