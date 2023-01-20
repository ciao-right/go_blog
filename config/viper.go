package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitViper() *viper.Viper {
	path, _ := os.Getwd()
	v := viper.New()
	v.AddConfigPath(path)
	fmt.Println(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		if err != nil {
			//viper.ConfigFileNotFoundError 未找到文件的错误
			panic(err)
		}
	}
	return v
}
