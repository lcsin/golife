package diy_encoder

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 参考Gin日志风格，定制Encoder，将日志的各个字段通过“|”进行分隔
func getGinStyleEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	// 自定义日期格式
	config.EncodeTime = zapcore.TimeEncoderOfLayout("[ZAP] 2006/01/02 - 15:04:05")
	// 日志级别以大写的方式写入
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// 完整路径输出调用文件
	config.EncodeCaller = zapcore.FullCallerEncoder
	// 分隔符
	config.ConsoleSeparator = " | "
	return zapcore.NewConsoleEncoder(config)
}
