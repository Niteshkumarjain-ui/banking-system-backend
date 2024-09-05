package util

import (
	"go.uber.org/zap"
)

var zap_sugared_logger *zap.SugaredLogger

func GetLogger() (logger *zap.SugaredLogger) {
	logger = zap_sugared_logger
	return
}

func init() {
	logger_config := zap.NewProductionConfig()

	switch _LEVEL_ := Configuration.Log.Level; _LEVEL_ {
	case "DEBUG":
		logger_config.Level.SetLevel(zap.DebugLevel)
	case "INFO":
		logger_config.Level.SetLevel(zap.InfoLevel)
	case "WARN":
		logger_config.Level.SetLevel(zap.WarnLevel)
	case "ERROR":
		logger_config.Level.SetLevel(zap.ErrorLevel)
	default:
		logger_config.Level.SetLevel(zap.DebugLevel)
	}

	logger_config.EncoderConfig.FunctionKey = "method"

	logger, _ := logger_config.Build()

	sugar_logger := logger.Sugar()
	zap_sugared_logger = sugar_logger

	defer logger.Sync()
}
