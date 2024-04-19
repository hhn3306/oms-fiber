package initialize

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"oms-fiber/global"
	"oms-fiber/middleware"
	"oms-fiber/router"
)

func Routers() *fiber.App {
	Router := fiber.New()
	Router.Use(middleware.DefaultLogger())
	InstallPlugin(Router)
	systemRouter := router.AllGroupRouterApp
	fmt.Println(systemRouter)

	Router.Get("/", func(c *fiber.Ctx) error {
		//return errors.New("mdfuck")
		//zap.L().Error("mdfk")
		global.Log.Info("grd")
		return c.SendString("Hello, World!")
	})

	Router.Get("/123", func(c *fiber.Ctx) error {
		return errors.New("mdfuck")
		//return c.SendString("Hello, World!")
	})
	return Router
}
