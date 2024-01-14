package yandexcloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetIAMToken(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err.Error())
	}
	result := GetIAMToken(os.Getenv("OAUTH_TOKEN"))
	fmt.Println(result.IAMToken)
}
