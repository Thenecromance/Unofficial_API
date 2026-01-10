package postmanconverter

import (
	"strings"

	"github.com/Thenecromance/BlizzardAPI/tools/updater"
)

func defaultInfo(info *Info) {
	info.PostmanId = "ad457392-feea-4932-be28-0d066949a706"
	info.Name = "Blizzard API"
	info.Schema = "https://schema.getpostman.com/json/collection/v2.0.0/collection.json"
	info.ExporterId = "33860151"
}
func defaultAuth(auth *Auth) {
	auth.Type = "oauth2"
	auth.Oauth2.TokenName = "BlizzardAPIToken"
	auth.Oauth2.AccessTokenUrl = "https://us.battle.net/oauth/token"
	auth.Oauth2.GrantType = "client_credentials"
	auth.Oauth2.AddTokenTo = "header"
	auth.Oauth2.AuthUrl = "https://us.battle.net/oauth/token"
	auth.Oauth2.ClientSecret = "INPUT YOUR SECRET HERE"
	auth.Oauth2.ClientId = "INPUT YOUR ID HERE"

}

func defaultEvent(events *[]Event) {
	var event Event
	event.Listen = "prerequest"
	event.Script.Type = "text/javascript"
	event.Script.Exec = []string{""}
	*events = append(*events, event)
}

func GeneratePostManCollection(group []*updater.ApiGroup) Content {
	var config Content

	defaultInfo(&config.Info)
	defaultAuth(&config.Auth)
	defaultEvent(&config.Events)
	temp := make(map[string]*Item)
	for _, g := range group {
		subItem := GenerateItems(g)
		//temp[subItem.Name] =
		if temp[subItem.Name] == nil {
			temp[subItem.Name] = &subItem
		} else {
			temp[subItem.Name].SubItems = append(temp[subItem.Name].SubItems, subItem.SubItems...) /*= append(temp[subItem.Name], subItem)*/
		}
	}

	for _, v := range temp {
		config.Items = append(config.Items, *v)
	}

	//config.Items
	config.Variable = []Variable{}
	return config
}

func GenerateItems(group *updater.ApiGroup) Item {
	return GeneratePostManItems(group)
}

func GeneratePostManItems(group *updater.ApiGroup) Item {
	var item Item
	item.Name = group.Game
	for _, api := range group.Methods {
		subItem := GeneratePostManSubItems(api)
		item.SubItems = append(item.SubItems, subItem)
	}
	return item
}

func GeneratePostManSubItems(api *updater.Api) SubItem {
	var item SubItem
	item.Name = api.Name

	var subItemInfo SubItemInfo
	subItemInfo.Name = api.Name
	subItemInfo.Request.Method = api.Method
	subItemInfo.Request.Header = []interface{}{}
	subItemInfo.Response = []interface{}{}

	subItemInfo.Request.Url.Query = []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{}
	for _, param := range api.Params {
		if !param.IsBindingUri {
			var queryParam struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}
			queryParam.Key = param.SourceName
			queryParam.Value = updater.AnyToString(param.DefaultValue)
			subItemInfo.Request.Url.Query = append(subItemInfo.Request.Url.Query, queryParam)
		}
	}
	// URL
	var urlParts []string
	urlParts = append(urlParts, "{{base_url}}")
	var convertedPath string
	convertedPath = api.Path
	convertedPath = strings.ReplaceAll(convertedPath, "{", ":")
	convertedPath = strings.ReplaceAll(convertedPath, "}", "")
	urlParts = append(urlParts, strings.Split(convertedPath, "/")...)

	subItemInfo.Request.Url.Raw = "{{base_url}}" + convertedPath
	subItemInfo.Request.Url.Host = []string{"{{base_url}}"}
	subItemInfo.Request.Url.Path = urlParts[1:]

	item.Item = append(item.Item, subItemInfo)
	return item

}
