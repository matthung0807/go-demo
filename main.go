package main

import (
	"go.uber.org/zap"
)

func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	sugar.Debug("debug message")
	sugar.Info("info message")
	sugar.Error("error message")
	sugar.Warn("warn message")
	// sugar.Panic("panic message")
	sugar.Fatal("fatal message")
}
