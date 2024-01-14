package yandexcloud

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type IAMTokenRequest struct {
	YandexPassportOautToken string `json:"yandexPassportOauthToken"`
}

type IAMTokenResponse struct {
	IAMToken  string    `json:"iamToken"`
	ExpiresAt time.Time `json:"expiresAt"`
}

const Endpoint = "https://iam.api.cloud.yandex.net/iam/v1/tokens"

func GetIAMToken(yandexOauthKey string) IAMTokenResponse {
	var i IAMTokenRequest
	i.YandexPassportOautToken = yandexOauthKey
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	Body := bytes.NewBuffer(data)
	resp, err := http.Post(Endpoint, "application/json", Body)
	if err != nil {
		panic(err)
	}
	responceData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var respStruct IAMTokenResponse
	err = json.Unmarshal(responceData, &respStruct)
	if err != nil {
		panic(err)
	}
	return respStruct
}
