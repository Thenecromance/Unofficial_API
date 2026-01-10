package main

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/Thenecromance/BlizzardAPI/bridge/log"
	"github.com/Thenecromance/BlizzardAPI/tools/postmanconverter"
	"github.com/Thenecromance/BlizzardAPI/tools/updater"
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

	LocalPath   string
	PostManFile string
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
			t.Result = updater.RequestApiList(t.Game, t.URL, t.Category)
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
		if task.Enabled && task.Result != nil {
			collection = append(collection, task.Result...)
		}

		outPath := "../api/" + task.Game + "/" + task.Category + "/"

		if f.Api && task.Enabled {
			log.Infof("Generating %s API functions...", task.Game)
			updater.GenerateApi(task.Game, outPath, task.Result)
		}

		if f.Model && task.Enabled {
			log.Infof("Generating %s Data Models...", task.Game)
			updater.GenerateModels(task.Game, outPath, task.Result)
		}
	}

	if f.Router {
		log.Info("Generating Router Mappings...")
		updater.GenerateRouters("routers", "../routers/", collection)
	}

	if len(f.LocalPath) > 0 {
		log.Infof("Storing API info to %s ...", f.LocalPath)
		buf, _ := json.MarshalIndent(collection, "", "  ")
		os.WriteFile(f.LocalPath, buf, 0644)
	}

	if len(f.PostManFile) > 0 {
		log.Infof("Generating PostMan Collection file %s ...", f.PostManFile)
		result := postmanconverter.GeneratePostManCollection(collection)
		jStr, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			panic(err)
		}
		os.WriteFile(f.PostManFile, jStr, 0644)
	}

}

func main() {
	updateApi(&Fields{
		Api:    false, // Generate API functions
		Model:  false, // Generate Data Models
		Router: false, // Generate Router Mappings (so far only support gin)

		Wow:        true, // Generate WoW Retail APIs
		Classic:    true, // Generate WoW Classic APIs
		D3:         true, // Generate Diablo 3 APIs
		HeartStone: true, // Generate HearthStone APIs
		SC2:        true, // Generate StarCraft II APIs

		LocalPath:   "./api_collection.json", // store all API info to a local file, if it is empty, app will not store these datas
		PostManFile: "./pm.auto_gen.json",    // Generate PostMan Collection file, if it is empty, app will not generate this file
	})
}
