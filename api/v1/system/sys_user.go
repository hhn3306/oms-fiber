package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"oms-fiber/global"
	"oms-fiber/model/common/response"
	"oms-fiber/model/system"
	systemReq "oms-fiber/model/system/request"
)

//var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b BaseApi) Login(c *fiber.Ctx) error {
	var l systemReq.Login
	if err := c.BodyParser(&l); err != nil {
		response.FailWithMessage(err.Error(), c)
		return nil
	}

	key := c.IP()
	if err := global.Validate.Struct(l); err != nil {
		zap.L().Info("1")
		response.FailWithMessage(err.Error(), c)
		return nil
	}

	//// 判断验证码是否开启
	//openCaptcha := global.Config.Captcha.OpenCaptcha               // 是否开启防爆次数
	//openCaptchaTimeOut := global.Config.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	//openCaptcha := 0        // 是否开启防爆次数
	//openCaptchaTimeOut := 2 // 缓存超时时间
	//v, ok := global.BlackCache.Get(key)
	//if !ok {
	//	global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	//}

	//var oc bool = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v)
	//if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.Log.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return nil
	}
	if user.Enable != 1 {
		global.Log.Error("登陆失败! 用户被禁止登录!")
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户被禁止登录", c)
		return nil
	}
	//b.TokenNext(c, *user)
	return nil
	//}
	//return nil
}
