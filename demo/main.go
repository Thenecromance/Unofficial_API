package main

import (
	"github.com/Thenecromance/BlizzardAPI/app"
	_ "github.com/Thenecromance/BlizzardAPI/routers"
)

func main() {
	app.Instance().Run()

}
