package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"oms-fiber/global"
	"oms-fiber/model/common/response"
	"strings"
)

type CustomClaims struct {
	jwt.MapClaims
}

func CustomJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(global.Config.JWT.SigningKey)},
		Claims:       &CustomClaims{},
		ErrorHandler: customJWTErrorHandler,
		//SuccessHandler: customSuccessHandler,  //有定制需求自己可以考虑搞搞。现有的就很丰满
		Filter: func(c *fiber.Ctx) bool {
			// Define the URLs that should be ignored by the JWT middleware
			return c.Path() == "/api/user/login"
		},
	})
}

func customJWTErrorHandler(c *fiber.Ctx, err error) error {
	// 这个错误类型多,可以尝试判断下类型了解下
	zap.L().Warn(err.Error())
	response.FailWithMessage(err.Error(), c)
	return nil
}

func extractJWTToken(authorizationHeader string) string {
	// 将 Authorization 头按空格拆分
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 {
		// 如果拆分后的部分不是两个，返回空字符串
		return ""
	}
	// 返回去掉 Bearer 字符串后的 Token 部分
	return parts[1]
}

//func customSuccessHandler(c *fiber.Ctx) error {
//	// 从请求头中获取 Token
//	tokenHeader := c.Get("Authorization")
//	token := extractJWTToken(tokenHeader)
//	j := utils.NewJWT()
//	// 解析 Token
//	claims, err := j.ParseToken(token)
//	if err != nil {
//		zap.L().Warn(err.Error())
//		response.FailWithMessage("token error", c)
//		return nil
//	}
//	// 将用户信息存储到上下文中
//	c.Locals("user", claims)
//	// 继续执行下一个处理程序
//	return c.Next()
//}
