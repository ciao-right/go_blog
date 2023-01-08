package config

import "github.com/spf13/viper"

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if err != nil {
			//viper.ConfigFileNotFoundError 未找到文件的错误
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic("配置文件未找到")
			} else {
				panic("未知错误")
			}
		}
	}
}
