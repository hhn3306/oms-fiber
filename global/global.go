package global

import (
	"github.com/bwmarrin/snowflake"
	v "github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"oms-fiber/config"
)

var (
	Config      config.Server
	Log         *zap.Logger
	Validate    *v.Validate
	BlackCache  local_cache.Cache
	Redis       *redis.Client
	SnowflakeID *snowflake.Node
	Gorm        *gorm.DB
	// ConcurrencyControl 用于在并发环境中控制对某个操作的并发执行
	ConcurrencyControl = &singleflight.Group{}
)
