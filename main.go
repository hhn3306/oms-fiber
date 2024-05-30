package main

import (
	"go.uber.org/zap"
	"oms-fiber/core"
	"oms-fiber/global"
	"oms-fiber/initialize"
)

func main() {
	core.Viper()            // 初始化Viper Config
	global.Log = core.Zap() // 初始化日志
	zap.ReplaceGlobals(global.Log)
	initialize.OtherInit()
	global.Gorm = initialize.Gorm()
	if global.Gorm != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.Gorm.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
