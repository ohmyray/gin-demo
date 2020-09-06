package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig()  {

	dirPath, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dirPath + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
