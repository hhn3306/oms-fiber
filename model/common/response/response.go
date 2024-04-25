package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 3000
	SUCCESS = 2000
)

func Result(code int, data interface{}, msg string, c *fiber.Ctx) {
	c.Status(http.StatusOK).JSON(Response{code, data, msg})
}

func Ok(c *fiber.Ctx) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *fiber.Ctx) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *fiber.Ctx) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *fiber.Ctx) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *fiber.Ctx) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(ERROR, data, message, c)
}
