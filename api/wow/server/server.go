package server

import (
	"Unofficial_API/utils"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type UniversalServerStatus struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	StatusName     string `json:"status_name"`
	StatusType     string `json:"status_type"`
	PopulationName string `json:"population_name"`
	PopulationType string `json:"population_type"`
	Category       string `json:"category"`
	Locale         string `json:"locale"`
	Timezone       string `json:"timezone"`
	Region         string `json:"region"`
	TypeName       string `json:"type_name"`
	TypeType       string `json:"type_type"`
}
type Status struct {
	Data struct {
		Realms []struct {
			Name            string      `json:"name"`
			Slug            string      `json:"slug"`
			Locale          string      `json:"locale"`
			Timezone        string      `json:"timezone"`
			Online          bool        `json:"online"`
			Category        string      `json:"category"`
			RealmLockStatus interface{} `json:"realmLockStatus"`
			Type            struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
				Enum string `json:"enum"`
			} `json:"type"`
			Population struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
				Enum string `json:"enum"`
			} `json:"population"`
		} `json:"Realms"`
	} `json:"data"`
}

type CNServerStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Id             int    `json:"id"`
			Name           string `json:"name"`
			Slug           string `json:"slug"`
			StatusName     string `json:"status_name"`
			StatusType     string `json:"status_type"`
			PopulationName string `json:"population_name"`
			PopulationType string `json:"population_type"`
			Category       string `json:"category"`
			Locale         string `json:"locale"`
			Timezone       string `json:"timezone"`
			Region         string `json:"region"`
			TypeName       string `json:"type_name"`
			TypeType       string `json:"type_type"`
		} `json:"List"`
	} `json:"data"`
}

func UpdateAllServerStatus() (string, error) {
	return "", nil
	//https://webapi.blizzard.cn/wow-armory-server/api/server_status?server_type=wow_mainline
}

func UpdateCNServerStatus() ([]UniversalServerStatus, error) {
	client := utils.NewRequest()

	resp, err := client.GET("https://webapi.blizzard.cn/wow-armory-server/api/server_status?server_type=wow_mainline")
	if err != nil {
		return nil, err
	}
	serverStatus := &CNServerStatus{}
	json.Unmarshal([]byte(resp), serverStatus)

	var result []UniversalServerStatus
	for _, v := range serverStatus.Data.List {
		result = append(result, UniversalServerStatus{
			Id:             v.Id,
			Name:           v.Name,
			Slug:           v.Slug,
			StatusName:     v.StatusName,
			StatusType:     v.StatusType,
			PopulationName: v.PopulationName,
		})
	}
	return result, nil

	//return client.GET("https://webapi.blizzard.cn/wow-armory-server/api/server_status?server_type=wow_mainline")
}
func UpdateTWServerStatus() ([]UniversalServerStatus, error) {
	resp, err := doRequest("tw")
	if err != nil {
		return nil, err
	}
	serverStatus := &Status{}
	err = json.Unmarshal([]byte(resp), serverStatus)
	if err != nil {
		return nil, err
	}

	var result []UniversalServerStatus
	for _, v := range serverStatus.Data.Realms {
		result = append(result, UniversalServerStatus{
			Id:         0,
			Name:       v.Name,
			Slug:       v.Slug,
			StatusName: v.Type.Name,
			StatusType: func() string {
				if v.Online {
					return "UP"
				}
				return "DOWN"
			}(),
			PopulationName: v.Population.Name,
			PopulationType: v.Population.Enum,
			Category:       v.Category,
			Locale:         v.Locale,
			Timezone:       v.Timezone,
			Region:         "tw",
			TypeName:       v.Type.Name,
			TypeType:       v.Type.Enum,
		})
	}

	return result, nil
}
func UpdateKRServerStatus() ([]UniversalServerStatus, error) {
	resp, err := doRequest("kr")
	if err != nil {
		return nil, err
	}
	serverStatus := &Status{}
	err = json.Unmarshal([]byte(resp), serverStatus)
	if err != nil {
		return nil, err
	}

	var result []UniversalServerStatus
	for _, v := range serverStatus.Data.Realms {
		result = append(result, UniversalServerStatus{
			Id:         0,
			Name:       v.Name,
			Slug:       v.Slug,
			StatusName: v.Type.Name,
			StatusType: func() string {
				if v.Online {
					return "UP"
				}
				return "DOWN"
			}(),
			PopulationName: v.Population.Name,
			PopulationType: v.Population.Enum,
			Category:       v.Category,
			Locale:         v.Locale,
			Timezone:       v.Timezone,
			Region:         "kr",
			TypeName:       v.Type.Name,
			TypeType:       v.Type.Enum,
		})
	}

	return result, nil
}

func UpdateEUServerStatus() ([]UniversalServerStatus, error) {
	return doRequest("eu")
}
func UpdateUSServerStatus() ([]UniversalServerStatus, error) {
	return doRequest("us")
}

func doRequest(region string) (string, error) {
	client := utils.NewRequest()
	body := fmt.Sprintf(`{"operationName":"GetRealmStatusData","variables":{"input":{"compoundRegionGameVersionSlug":"%s"}},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"b37e546366a58e211e922b8c96cd1ff74249f564a49029cc9737fef3300ff175"}}}`, region)
	return client.POSTWithBody("https://worldofwarcraft.blizzard.com/en-us/graphql", NewStringReader(body))
}
func NewStringReader(s string) io.Reader {
	return strings.NewReader(s)
}
