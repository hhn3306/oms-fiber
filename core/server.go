package core

import (
	"fmt"
	"oms-fiber/global"
	"oms-fiber/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
	//	// 初始化redis服务
	//	initialize.Redis()
	//}

	//// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	//address := fmt.Sprintf(":%d", global.Config.System.Addr)
	address := fmt.Sprintf(":%d", 3000)

	//	fmt.Printf(`
	//	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	//	默认前端文件运行地址:http://127.0.0.1:8080
	//`, address)
	global.Log.Error(Router.Listen(address).Error())
}
