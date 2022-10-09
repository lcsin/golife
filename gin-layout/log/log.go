package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golife/gin-layout/config"
	"github.com/golife/gin-layout/pkg"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ginStyle     = "gin"
	consoleStyle = "console"
	jsonStyle    = "json"
)

type Log func(args ...interface{})
type Logf func(template string, args ...interface{})

var zapLog *zap.SugaredLogger
var (
	Debug  Log
	Debugf Logf

	Info  Log
	Infof Logf

	Warn  Log
	Warnf Logf

	Error  Log
	Errorf Logf

	Fatal  Log
	Fatalf Logf
)

func Init() {
	var cores []zapcore.Core
	// 标准输出
	encoder := getEncoder()
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	cores = append(cores, c1)
	// error日志单独输出到文件
	if config.App().Zap.Error {
		ts := time.Now().Unix()
		file, err := os.Create(fmt.Sprintf("./log/%d.error.log", ts))
		if err != nil {
			log.Fatal(err)
		}
		c2 := zapcore.NewCore(encoder, zapcore.AddSync(file), zapcore.ErrorLevel)
		cores = append(cores, c2)
	}
	// 日志切割
	if config.App().Zap.Cut {
		lumberJackLogger := getLumberJackLogger()
		c3 := zapcore.NewCore(encoder, lumberJackLogger, zapcore.InfoLevel)
		cores = append(cores, c3)
	}
	// 初始化日志输出
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())
	zapLog = logger.Sugar()
	initLogs()
}

func initLogs() {
	Debug = zapLog.Debug
	Debugf = zapLog.Debugf

	Info = zapLog.Info
	Infof = zapLog.Infof

	Warn = zapLog.Warn
	Warnf = zapLog.Warnf

	Error = zapLog.Error
	Errorf = zapLog.Errorf

	Fatal = zapLog.Fatal
	Fatalf = zapLog.Fatalf
}

func getConsoleStyleEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	// 自定义日期格式
	config.EncodeTime = zapcore.TimeEncoderOfLayout(pkg.DefaultTimeFormat)
	// 日志级别以大写的方式写入
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// 完整路径输出调用文件
	config.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(config)
}

func getJSONStyleEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	// 自定义日期格式
	config.EncodeTime = zapcore.TimeEncoderOfLayout(pkg.DefaultTimeFormat)
	// 日志级别以大写的方式写入
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// 短路径输出调用文件
	config.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(config)
}

func getGinStyleEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	// 自定义日期格式
	config.EncodeTime = zapcore.TimeEncoderOfLayout("[ZAP] 2006/01/02 - 15:04:05")
	// 日志级别以大写的方式写入
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// 完成路径输出调用文件
	config.EncodeCaller = zapcore.FullCallerEncoder
	// 分隔符
	config.ConsoleSeparator = " | "
	return zapcore.NewConsoleEncoder(config)
}

func getEncoder() zapcore.Encoder {
	switch config.App().Zap.Style {
	case ginStyle:
		return getGinStyleEncoder()
	case consoleStyle:
		return getConsoleStyleEncoder()
	case jsonStyle:
		return getJSONStyleEncoder()
	default:
		log.Fatalf("this log style is not supported: %v", config.App().Zap.Style)
	}
	return nil
}

// 日志切割
func getLumberJackLogger() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.App().Zap.Path, // 日志文件的位置
		MaxSize:    1,                     // 在进行切割之前，日志文件的最大大小，单位MB
		MaxBackups: 5,                     // 保留旧文件的最大个数
		MaxAge:     30,                    // 保留旧文件的最大天数
		Compress:   false,                 // 是否压缩或归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
