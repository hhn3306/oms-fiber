package initialize

import (
	"github.com/gofiber/fiber/v2"
	"oms-fiber/global"
	"oms-fiber/middleware"
	"oms-fiber/router"
)

func Routers() *fiber.App {
	Router := fiber.New()
	Router.Use(middleware.DefaultLogger())
	//InstallPlugin(Router)
	systemRouter := router.AllGroupRouterApp.System

	//PublicGroup := Router.Group(global.Config.System.RouterPrefix).Use(middleware.CustomJWTMiddleware())
	PublicGroup := Router.Group(global.Config.System.RouterPrefix).Use(middleware.CustomJWTMiddleware())

	systemRouter.InitBaseRouter(PublicGroup)

	return Router
}
