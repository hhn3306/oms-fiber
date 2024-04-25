package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "oms-fiber/api/v1"
)

type BaseRouter struct{}

//func (s *BaseRouter) InitBaseRouter(Router *fiber.App) {
//	userRouter := Router.Group("user").Use()
//	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
//	{
//		userRouter.Post("login", baseApi.Login)
//	}
//}

func (s *BaseRouter) InitBaseRouter(router fiber.Router) {
	userRouter := router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.Post("login", baseApi.Login)
	}
}
