package yandexcloud

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const CloudIDListEndpoint = "https://resource-manager.api.cloud.yandex.net/resource-manager/v1/clouds"

type CloudListCommonResponce struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OrganizationID string    `json:"organizationId"`
	Labels         string    `json:"labels"`
}

type CloudListCommonRequest struct {
	PageSize       uint16 `json:"pageSize"`
	PageToken      string `json:"pageToken"`
	Filter         string `json:"filter"`
	OrganizationID string `json:"organizationId"`
}

type CloudListResponce struct {
	Clouds        []*CloudListCommonResponce `json:"clouds"`
	NextPageToken string                     `json:"nextPageToken"`
}

func GetCloudsList(iamToken string) (CloudListResponce, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", CloudIDListEndpoint, nil)
	if err != nil {
		return CloudListResponce{}, err
	}
	request.Header.Add("Authorization", GetYandexCloudHeader(iamToken))
	respone, err := client.Do(request)
	if err != nil {
		return CloudListResponce{}, err
	}
	responceData, err := io.ReadAll(respone.Body)
	if err != nil {
		return CloudListResponce{}, err
	}
	var result CloudListResponce
	err = json.Unmarshal(responceData, &result)
	if err != nil {
		return CloudListResponce{}, err
	}
	return result, nil
}
