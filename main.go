package main

import (
	"Unofficial_API/api/wow/DataService/Achievement"
	"Unofficial_API/global"
	"context"
	"encoding/json"
)

func main() {
	achievement, err := wowRetail_Achievement.AchievementCategoriesIndex(context.Background(), &wowRetail_Achievement.AchievementCategoriesIndexFields{
		Namespace: global.Static(),
		Locale:    global.Locale(),
	})
	if err != nil {
		return
	}
	resp, _ := json.MarshalIndent(achievement, "", "\t")
	println(string(resp))
}
