package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"oms-fiber/global"
	"oms-fiber/model/system"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.Gorm
	err := db.AutoMigrate(
		// 系统模块表
		system.SysUser{},
	)
	if err != nil {
		zap.L().Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	//global.Log.Info("register table success")
	zap.L().Info("register table success")
}
