package logger

import (
	"go.uber.org/zap"
)

// ProvideLogger provides a zap logger
// TODO: later replace with logrus
func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
	//logrus.SetFormatter(new(logrus.JSONFormatter))
}

var Options = ProvideLogger
