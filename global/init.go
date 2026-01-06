package global

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("universal_config")
	viper.SetConfigType("yaml")

	//viper.AddConfigPath("/etc/unofficial_api/")
	//viper.AddConfigPath("$HOME/.unofficial_api")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {

		err := viper.SafeWriteConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Config file loaded successfully")
		cookie := GetCookie()
		fmt.Println("Cookie:", cookie[:4]+"...."+cookie[len(cookie)-4:])

	}

}
