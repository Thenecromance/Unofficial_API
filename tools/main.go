package main

import (
	"Unofficial_API/tools/updater"
	"sync"
)

func updateApi() {
	var wg sync.WaitGroup
	wg.Add(5)

	// retail version
	/*go*/
	func() {
		defer wg.Done()
		//updater.ParseTemplate("wowRetail", "./api/wow/DataService/", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/game-data-apis.json")
		updater.ParseTemplate("wowRetail", "./api/wow/ProfileService/", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/profile-apis.json")
	}()

	// Wow Classic version
	go func() {
		defer wg.Done()
		//updater.ParseTemplate("wowClassic", "./api/wowClassic/DataService/", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/game-data-apis.json")
		updater.ParseTemplate("wowClassic", "./api/wowClassic/ProfileService/", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/profile-apis.json")
	}()
	// D3 support
	go func() {
		defer wg.Done()
		// CN data not supported
		//parseTemplate("D3", "./api/D3/Community/CN/", "https://develop.battle.net/api/pages/content/documentation/diablo-3/community-apis-cn.json")
		updater.ParseTemplate("D3", "./api/D3/Community/", "https://develop.battle.net/api/pages/content/documentation/diablo-3/community-apis.json")
		updater.ParseTemplate("D3", "./api/D3/DataService/", "https://develop.battle.net/api/pages/content/documentation/diablo-3/game-data-apis.json")
	}()
	//HeartStone support
	go func() {
		defer wg.Done()

		updater.ParseTemplate("HeartStone", "./api/HeartStone/DataService/", "https://develop.battle.net/api/pages/content/documentation/hearthstone/game-data-apis.json")
	}()
	go func() {
		defer wg.Done()
		// CN data not supported
		//parseTemplate("SC", "./api/StarCraftII/Community/CN/", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis-cn.json")
		updater.ParseTemplate("SC", "./api/StarCraftII/Community/", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis.json")
		updater.ParseTemplate("SC", "./api/StarCraftII/DataService/", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/game-data-apis.json")
	}()

	wg.Wait()
}

func main() {
	updateApi()
}
