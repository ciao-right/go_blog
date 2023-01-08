package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	if db, err := gorm.Open(mysql.Open(getDbConfig()), &gorm.Config{}); err != nil {
		return db
	} else {
		panic("连接失败")
	}
}

func getDbConfig() string {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	address := viper.GetString("mysql.password")
	port := viper.GetString("mysql.port")
	databaseName := viper.GetString("mysql.database_name")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, databaseName)
}

func RegisterTables(db *gorm.DB) {
	// 注册表
	db.AutoMigrate()
}
