package main

import (
	"Unofficial_API/tools/updater"
	"encoding/json"
	"os"
	"sync"
)

type Fields struct {
	Api    bool
	Model  bool
	Router bool

	Wow        bool
	Classic    bool
	D3         bool
	HeartStone bool
	SC2        bool

	LocalPath string
}

func updateApi(f *Fields) {

	type apiTask struct {
		Game     string
		Category string
		URL      string
		Enabled  bool
		Result   []*updater.ApiGroup
	}

	tasks := []*apiTask{
		// Wow Retail
		{"wow", "DataService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/game-data-apis.json", f.Wow, nil},
		{"wow", "ProfileService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/profile-apis.json", f.Wow, nil},
		// Wow Classic
		{"wowClassic", "DataService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/game-data-apis.json", f.Classic, nil},
		{"wowClassic", "ProfileService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/profile-apis.json", f.Classic, nil},
		// Diablo 3
		{"D3", "Community", "https://develop.battle.net/api/pages/content/documentation/diablo-3/community-apis.json", f.D3, nil},
		{"D3", "DataService", "https://develop.battle.net/api/pages/content/documentation/diablo-3/game-data-apis.json", f.D3, nil},
		// Hearthstone
		{"HeartStone", "DataService", "https://develop.battle.net/api/pages/content/documentation/hearthstone/game-data-apis.json", f.HeartStone, nil},
		// StarCraft II
		{"StarCraftII", "Community", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis.json", f.SC2, nil},
		{"StarCraftII", "DataService", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/game-data-apis.json", f.SC2, nil},
	}

	var wg sync.WaitGroup
	for i := range tasks {
		wg.Add(1)
		go func(t *apiTask) {
			defer wg.Done()
			t.Result = updater.RequestApiList(t.Game, t.URL)
		}(tasks[i])

	}
	wg.Wait()

	for i := range tasks {
		for j := range tasks[i].Result {
			tasks[i].Result[j].Fixed()
		}
	}

	collection := make([]*updater.ApiGroup, 0)

	for _, task := range tasks {
		if task.Result != nil {
			collection = append(collection, task.Result...)
		}

		outPath := "./api/" + task.Game + "/" + task.Category + "/"

		if f.Api && task.Enabled {
			updater.GenerateApi(task.Game, outPath, task.Result)
		}

		if f.Model && task.Enabled {
			updater.GenerateModels(task.Game, outPath, task.Result)
		}
	}

	if f.Router {
		updater.GenerateRouters("routers", "./routers/", collection)
	}

	if len(f.LocalPath) > 0 {
		buf, _ := json.MarshalIndent(collection, "", "  ")
		os.WriteFile(f.LocalPath, buf, 0644)
	}

}

func main() {
	updateApi(&Fields{
		Api:    true,
		Model:  true,
		Router: true,

		Wow:        true,
		Classic:    true,
		D3:         true,
		HeartStone: true,
		SC2:        true,

		LocalPath: "./api_collection.json",
	})
}
