package global

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("client_id", "")
	viper.SetDefault("client_secret", "")
}
func GetClientID() string {
	return viper.GetString("client_id")
}
func GetClientSecret() string {
	return viper.GetString("client_secret")
}
