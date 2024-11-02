package utils

import (
	config "github.com/hrabit64/sproutnote/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	SchedulerLoggerName = "scheduler"
	CilLoggerName       = "cil"
)

func GetLogger() (*zap.Logger, error) {
	loggerName := config.ProcessType

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout", "./logs/" + loggerName + ".log"},
		ErrorOutputPaths: []string{"stderr", "./logs/" + loggerName + "_error.log"},
	}

	return config.Build()

}
