package config

import (
	"fmt"
	"go_blog/common/global"
	"go_blog/model/request"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	if db, err := gorm.Open(mysql.Open(getDbConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err == nil {
		return db
	} else {
		fmt.Println(err)
		panic("连接失败")
	}
}

func getDbConfig() string {
	username := global.GlobalViper.GetString("mysql.username")
	password := global.GlobalViper.GetString("mysql.password")
	address := global.GlobalViper.GetString("mysql.address")
	port := global.GlobalViper.GetString("mysql.port")
	databaseName := global.GlobalViper.GetString("mysql.database_name")
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, databaseName))
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, databaseName)
}

func RegisterTables(db *gorm.DB) {
	// 注册表
	db.AutoMigrate(request.User{})
	//db.AutoMigrate(request.GoodsClassification{})
}
