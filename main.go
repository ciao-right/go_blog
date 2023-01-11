package main

import (
	"go_blog/common/global"
	"go_blog/config"
)

func main() {
	//初始化 viper
	config.InitViper()
	//初始化数据库
	global.GLOBAL_DB = config.InitDb()
	if global.GLOBAL_DB != nil {
		config.RegisterTables(global.GLOBAL_DB)
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}
	r := config.Routers()
	r.Run()
}
