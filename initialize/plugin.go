package initialize

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func InstallPlugin(Router *fiber.App) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	//PrivateGroup.Use().Use()
}
