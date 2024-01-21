package yandexcloud

import "fmt"

type YandexCloudHeader struct {
	Authorization string `json:"Authorization"`
}

func GetYandexCloudHeader(iamToken string) string {
	return fmt.Sprintf("Bearer %s", iamToken)
}
