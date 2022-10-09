package diy_logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 将error日志单独输出到文件，其它级别的日志输出到控制台
func getRecordErrorLogger() *zap.Logger {
	// 使用zap提供的编码器以及配置
	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	// 创建第一个zapcore.Core，将日志写到控制台
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	// 创建第二个zapcore.Core，将error日志写到文件
	file, _ := os.Create("./error.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(file), zapcore.ErrorLevel)
	// 用c1和c2两个zapcore.Core处理日志
	core := zapcore.NewTee(c1, c2)
	// zap.AddCaller()方法打印日志的文件调用路径
	return zap.New(core, zap.AddCaller())
}
