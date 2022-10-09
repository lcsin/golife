package diy_encoder

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestEncoder(t *testing.T) {
	encoder := getGinStyleEncoder()
	core := zapcore.NewCore(encoder, os.Stdout, zapcore.ErrorLevel)
	sugar := zap.New(core, zap.AddCaller()).Sugar()
	sugar.Error("test custom encoder")
}
