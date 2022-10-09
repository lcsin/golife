package join_lumberjack

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

func getLumberJackLogger() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap.log", // 日志文件的位置
		MaxSize:    1,           // 在进行切割之前，日志文件的最大大小，单位MB
		MaxBackups: 5,           // 保留旧文件的最大个数
		MaxAge:     30,          // 保留旧文件的最大天数
		Compress:   false,       // 是否压缩或归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
