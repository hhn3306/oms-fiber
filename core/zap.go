package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"oms-fiber/core/internal"
	"oms-fiber/global"
	"oms-fiber/utils"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

//
//
//package initialize
//
//import (
//"fmt"
//"go.uber.org/zap"
//"go.uber.org/zap/zapcore"
//"time"
//)
//
//var logger *zap.Logger
//
//func InitLogger() {
//	// 创建 Zap 日志记录器配置
//	zapConfig := zap.NewDevelopmentConfig()
//	// 添加 caller
//	//zapConfig.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
//	zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
//
//	// 自定义日志时间格式
//	zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
//	}
//
//	// 根据配置文件设置日志级别
//	//zapConfig.Level = zap.NewAtomicLevelAt(StringToLevel(global.Config.Zap.Level))
//	zapConfig.Level = zap.NewAtomicLevelAt(StringToLevel("info"))
//
//	// 构建日志记录器实例
//	var err error
//	logger, err = zapConfig.Build()
//	if err != nil {
//		fmt.Printf("Failed to build logger: %v\n", err)
//		return
//	}
//}
//
//func GetLogger() *zap.Logger {
//	if logger == nil {
//		InitLogger()
//	}
//	return logger
//}
//
//type Level int8
//
//const (
//	DebugLevel Level = iota - 1
//	InfoLevel
//	WarnLevel
//	ErrorLevel
//	DPanicLevel
//	PanicLevel
//	FatalLevel
//)
//
//func StringToLevel(levelStr string) zapcore.Level {
//	switch levelStr {
//	case "debug":
//		return zapcore.DebugLevel
//	case "info":
//		return zapcore.InfoLevel
//	case "warn":
//		return zapcore.WarnLevel
//	case "error":
//		return zapcore.ErrorLevel
//	case "dpanic":
//		return zapcore.DPanicLevel
//	case "panic":
//		return zapcore.PanicLevel
//	case "fatal":
//		return zapcore.FatalLevel
//	default:
//		return zapcore.InfoLevel // 默认为 InfoLevel
//	}
//}
