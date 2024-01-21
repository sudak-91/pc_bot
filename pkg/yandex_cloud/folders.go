package yandexcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const FolderListEndpoint = "https://resource-manager.api.cloud.yandex.net/resource-manager/v1/folders"

type FoldersListRequset struct {
	CloudID   string `json:"cloudId"`
	PageSize  uint16 `json:"pageSize,omitempty"`
	PageToken string `json:"pageToken,omitempty"`
	Filter    string `json:"filter,omitempty"`
}

type FoldersListResponceCommon struct {
	Id          string    `json:"id"`
	CloudId     string    `json:"cloudId"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Labels      string    `json:"labels"`
	Status      string    `json:"status"`
}

type FoldersListResponce struct {
	List          []*FoldersListResponceCommon `json:"folders"`
	NextPageToken string                       `json:"nextPageToken"`
}

func GetFoldersList(cloudId string, iamToken string) (FoldersListResponce, error) {
	client := http.Client{}
	var reqListBody FoldersListRequset
	reqListBody.CloudID = cloudId

	data, err := json.Marshal(reqListBody)
	if err != nil {
		return FoldersListResponce{}, err
	}
	Body := bytes.NewBuffer(data)
	req, err := http.NewRequest("GET", FolderListEndpoint, Body)
	if err != nil {
		return FoldersListResponce{}, err
	}
	req.Header.Add("Authorization", GetYandexCloudHeader(iamToken))

	resp, err := client.Do(req)
	if err != nil {
		return FoldersListResponce{}, nil
	}
	responceData, err := io.ReadAll(resp.Body)
	if err != nil {
		return FoldersListResponce{}, err
	}
	fmt.Println(string(responceData))
	var result FoldersListResponce
	err = json.Unmarshal(responceData, &result)
	if err != nil {
		return FoldersListResponce{}, err
	}
	return result, nil
}
