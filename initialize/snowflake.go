package initialize

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"oms-fiber/global"
)

func SnowflakeNode() {
	snowflake.NodeBits = 6
	snowflake.StepBits = 8
	node, err := snowflake.NewNode(int64(global.Config.System.SnowflakeNode))
	if err != nil {
		zap.L().Error("雪花算法错误")
	}
	global.SnowflakeID = node
}
