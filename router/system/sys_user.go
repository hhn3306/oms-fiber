package system

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *fiber.App) {
	userRouter := Router.Group("user").Use()
	//todo
	fmt.Println(userRouter)
}
