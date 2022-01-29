package log

import "go.uber.org/zap"

func NewZapLogger() Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logger.Sync() }()
	return logger.Sugar()
}
