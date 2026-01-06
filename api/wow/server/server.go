package server

import (
	"Unofficial_API/bridge/client"
	"Unofficial_API/bridge/log"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
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
	serverList := make([][]UniversalServerStatus, 5)
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		log.Debug("Updating CN server status")
		defer log.Debug("Finished updating CN server status")
		defer wg.Done()
		var err error
		serverList[0], err = UpdateCNServerStatus()
		if err != nil {
			fmt.Println("Error updating CN server status:", err)
		}

	}()
	go func() {
		log.Debug("Updating TW server status")
		defer log.Debug("Finished updating TW server status")
		defer wg.Done()
		var err error
		serverList[1], err = UpdateTWServerStatus()
		if err != nil {
			fmt.Println("Error updating TW server status:", err)
		}
	}()
	go func() {
		log.Debug("Updating US server status")
		defer log.Debug("Finished updating US server status")
		defer wg.Done()
		var err error
		serverList[2], err = UpdateUSServerStatus()
		if err != nil {
			fmt.Println("Error updating US server status:", err)
		}
	}()
	go func() {
		log.Debug("Updating EU server status")
		defer log.Debug("Finished updating EU server status")
		defer wg.Done()
		var err error
		serverList[3], err = UpdateEUServerStatus()
		if err != nil {
			fmt.Println("Error updating EU server status:", err)
		}
	}()
	go func() {
		log.Debug("Updating KR server status")
		defer log.Debug("Finished updating KR server status")
		defer wg.Done()
		var err error
		serverList[4], err = UpdateKRServerStatus()
		if err != nil {
			fmt.Println("Error updating KR server status:", err)
		}

	}()
	wg.Wait()

	bytes, err := json.Marshal(serverList)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func convertStatusToUniversal(region string, status *Status) (result []UniversalServerStatus) {
	for _, v := range status.Data.Realms {
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
			Region:         region,
			TypeName:       v.Type.Name,
			TypeType:       v.Type.Enum,
		})
	}

	return result
}

func UpdateCNServerStatus() ([]UniversalServerStatus, error) {

	resp, err := client.GET("https://webapi.blizzard.cn/wow-armory-server/api/server_status?server_type=wow_mainline")
	if err != nil {
		return nil, err
	}
	serverStatus := &CNServerStatus{}
	err = json.Unmarshal([]byte(resp), serverStatus)
	if err != nil {
		return nil, err
	}
	var result []UniversalServerStatus
	for _, v := range serverStatus.Data.List {
		result = append(result, UniversalServerStatus{
			Id:             v.Id,
			Name:           v.Name,
			Slug:           v.Slug,
			StatusName:     v.StatusName,
			StatusType:     v.StatusType,
			PopulationName: v.PopulationName,
			PopulationType: v.PopulationType,
			Category:       v.Category,
			Locale:         v.Locale,
			Timezone:       v.Timezone,
			Region:         "cn",
			TypeName:       v.TypeName,
			TypeType:       v.TypeType,
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

	result := convertStatusToUniversal("kr", serverStatus)

	return result, nil
}

func UpdateEUServerStatus() ([]UniversalServerStatus, error) {
	resp, err := doRequest("eu")
	if err != nil {
		return nil, err
	}
	serverStatus := &Status{}
	err = json.Unmarshal([]byte(resp), serverStatus)
	if err != nil {
		return nil, err
	}

	result := convertStatusToUniversal("eu", serverStatus)

	return result, nil
}
func UpdateUSServerStatus() ([]UniversalServerStatus, error) {
	resp, err := doRequest("us")
	if err != nil {
		return nil, err
	}
	serverStatus := &Status{}
	err = json.Unmarshal([]byte(resp), serverStatus)
	if err != nil {
		return nil, err
	}

	result := convertStatusToUniversal("us", serverStatus)

	return result, nil
}

func doRequest(region string) (string, error) {

	body := fmt.Sprintf(`{"operationName":"GetRealmStatusData","variables":{"input":{"compoundRegionGameVersionSlug":"%s"}},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"b37e546366a58e211e922b8c96cd1ff74249f564a49029cc9737fef3300ff175"}}}`, region)
	return client.POSTWithBody("https://worldofwarcraft.blizzard.com/en-us/graphql", newStringReader(body))
}
func newStringReader(s string) io.Reader {
	return strings.NewReader(s)
}
