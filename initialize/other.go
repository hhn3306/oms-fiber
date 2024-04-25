package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"oms-fiber/global"
	"oms-fiber/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.Config.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.Config.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	// 初始化本地缓存
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	// 初始化参数效验函数 *validator.Validate
	global.Validate = validator.New(validator.WithRequiredStructEnabled())

	// 初始化雪花唯一ID
	SnowflakeNode()
}
