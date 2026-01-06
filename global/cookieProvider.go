package global

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("cookie", "")
}
func GetCookie() string {
	return viper.GetString("cookie")
}
