package pkg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func Init() {
	encoder := getCustomEncoder()
	core := zapcore.NewCore(encoder, os.Stdout, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zapLog = logger.Sugar()
	initLogs()
}

func getCustomEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.ConsoleSeparator = " / "
	return zapcore.NewConsoleEncoder(config)
}
