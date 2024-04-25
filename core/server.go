package core

import (
	"fmt"
	"oms-fiber/global"
	"oms-fiber/initialize"
)

func RunWindowsServer() {
	if global.Config.System.UseMultipoint || global.Config.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	//// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	//address := fmt.Sprintf(":%d", global.Config.System.Addr)
	address := fmt.Sprintf(":%d", 3000)

	global.Log.Error(Router.Listen(address).Error())
}
