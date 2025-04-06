package logger

import (
	"go.uber.org/zap"
)

func NewFileLogger(filePath string) *zap.SugaredLogger {
	return NewLogger(filePath)
}

func NewConsoleLogger() *zap.SugaredLogger {
	return NewLogger("stdout")
}

func NewFileLoggerWithConsole(filePath string) *zap.SugaredLogger {
	return NewLogger("stdout", filePath)
}

func NewLogger(filePath ...string) *zap.SugaredLogger {
	var config = zap.NewDevelopmentConfig()
	config.OutputPaths = filePath
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, _ := config.Build()
	return logger.Sugar()
}
