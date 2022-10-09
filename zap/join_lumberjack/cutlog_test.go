package join_lumberjack

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCutLog(t *testing.T) {
	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	writeSyncer := getLumberJackLogger()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	sugar := zap.New(core, zap.AddCaller()).Sugar()

	for i := 0; i < 10000; i++ {
		sugar.Debug("test log cut")
		sugar.Info("test log cut")
		sugar.Warn("test log cut")
		sugar.Error("test log cut")
	}
}
