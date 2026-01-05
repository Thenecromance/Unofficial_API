package main

import (
	"Unofficial_API/api/wow/server"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//res := wow.CNPlayerSummary(context.Background(), "黑色卷卷毛", "blanchard")

	/*res, err := wowRetail.BNetCharacterProfileSummary(context.Background(), "blanchard", "黑色卷卷毛")
	if err != nil {
		panic(err)
	}
	buff, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	os.WriteFile("out.json", buff, 0644)*/

	res, err := server.UpdateCNServerStatus()
	if err != nil {
		panic(err)
	}
	os.WriteFile("server_status.json", []byte(res), 0644)
}
