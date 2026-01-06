package main

import (
	"Unofficial_API/api/wow/ProfileService/CharacterProfile"
	"context"
	"encoding/json"
)

func main() {
	character, err := wow_CharacterProfile.CharacterProfileSummary(context.Background(), &wow_CharacterProfile.CharacterProfileSummaryFields{
		RealmSlug:     "ragnaros",
		CharacterName: "xhapira",
		Namespace:     "",
		Locale:        "zh_CN",
		ExtraFields:   nil,
		CN: /*&utils.CNRequestMethod{
			RealmSlug: "blanchard",
			Name:      "黑色卷卷毛",
		},*/nil,
	})
	if err != nil {
		return
	}
	resp, _ := json.MarshalIndent(character, "", "\t")
	println(string(resp))

	/*tpl, _ := uritemplates.Parse("/profile/wow/character/{realmSlug}/{characterName}")
	u3, _ := tpl.Expand(map[string]interface{}{"realmSlug": "shit", "characterName": "shit"})
	fmt.Println("uritemplates:", u3)*/

}
