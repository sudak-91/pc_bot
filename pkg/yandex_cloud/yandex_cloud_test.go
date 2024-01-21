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

	cloudsList, err := GetCloudsList(result.IAMToken)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%+v", cloudsList.Clouds[0].ID)

	foldersList, err := GetFoldersList(cloudsList.Clouds[0].ID, result.IAMToken)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%v", foldersList.List[0])

}
