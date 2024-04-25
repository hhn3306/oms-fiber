package middleware

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"oms-fiber/global"
	//"oms-fiber/global"
)

type log struct{}

func (l log) SetLoggerMiddleware() fiber.Handler {

	return fiberzap.New(fiberzap.Config{
		Logger: global.Log,
		// fiberzap.New查看有哪些
		Fields:   []string{"status", "ip", "url", "method", "latency"},
		Messages: []string{"ServerError", "ClientError", "success_response"},
		//FieldsFunc: func(c *fiber.Ctx) []zap.Field {
		//	// 获取调用者信息
		//	_, file, line, ok := runtime.Caller(8) // 这里根据具体情况调整层数
		//	if ok {
		//		return []zap.Field{zap.String("caller", file+":"+strconv.Itoa(line))}
		//	}
		//	return []zap.Field{zap.String("caller", "unknown")}
		//},
	})
}

func DefaultLogger() fiber.Handler {
	return log{}.SetLoggerMiddleware()
}
