package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	inline "github.com/sudak-91/telegrambotgo/telegram_api/inline_mode"
	"github.com/sudak-91/telegrambotgo/telegram_api/types"
)

func (t *TelegramUpdater) inlineQueryService(inlineQuery types.InlineQuery) ([]byte, error) {
	log.Println("INLINE QUERY SERVICE")
	var answer inline.AnswerInlineQuery
	answer.InlineQueryId = inlineQuery.ID
	answer.IsPersonal = true
	var article inline.InlineQueryResultArticle
	article.ID = uuid.New().String()
	article.Title = inlineQuery.Query
	article.Description = "New Test Inline Article"
	article.Type = "article"
	var messageConm inline.InputTextMessageContent
	messageConm.MessageText = "MessageText"

	article.InputMessageContent = messageConm

	answer.Results = append(answer.Results, article)
	data, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	log.Println(string(data))

	URL := fmt.Sprintf("https://api.telegram.org/bot%s/answerInlineQuery", os.Getenv("BOT_KEY"))
	log.Println(URL)
	Body := bytes.NewBuffer(data)
	resp, err := http.Post(URL, "application/json", Body)
	if err != nil {
		panic(err)

	}
	log.Println(resp)
	log.Println(resp.Header)
	log.Println(resp.Body)
	log.Println(resp.Request.Response)
	return data, nil
}
