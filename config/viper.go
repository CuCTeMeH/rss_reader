package config

import (
	"github.com/spf13/viper"
)

var Settings *viper.Viper

//Load viper instance to be accessible with the loaded .env config.
func ReadConfig() error {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	Settings = viper.GetViper()
	return nil
}
