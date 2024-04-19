package main

import (
	"go.uber.org/zap"
	"oms-fiber/core"
	"oms-fiber/global"
)

func main() {
	global.Log = core.Zap()
	zap.ReplaceGlobals(global.Log)

	core.RunWindowsServer()
}

//package main
//
//import (
//	"errors"
//	"github.com/gofiber/contrib/fiberzerolog"
//	"github.com/gofiber/fiber/v2"
//	"github.com/rs/zerolog"
//	"os"
//)
//
//func main() {
//	app := fiber.New()
//	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
//
//	app.Use(fiberzerolog.New(fiberzerolog.Config{
//		Logger: &logger,
//	}))
//
//	app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendString("Hello, World!")
//	})
//
//	app.Get("/1", func(c *fiber.Ctx) error {
//		return errors.New("mdfk")
//	})
//
//	if err := app.Listen(":3000"); err != nil {
//		logger.Fatal().Err(err).Msg("Fiber app error")
//	}
//}
