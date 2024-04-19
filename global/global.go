package global

import (
	"go.uber.org/zap"
	"oms-fiber/config"
)

var (
	Config config.Server
	Log    *zap.Logger
)
