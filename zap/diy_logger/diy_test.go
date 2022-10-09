package diy_logger

import "testing"

func TestErrorLogger(t *testing.T) {
	sugar := getRecordErrorLogger().Sugar()
	for i := 0; i < 10; i++ {
		sugar.Infof("test info log %d", i)
		sugar.Errorf("test error log %d", i)
	}
}
