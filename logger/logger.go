package logger

import (
	"github.com/Nixson/logNx"
	"gorm.io/gorm/logger"
)

type Writer struct {
	LogLevel logger.LogLevel
}

func (w *Writer) Write(p []byte) (n int, err error) {
	switch w.LogLevel {
	case logger.Silent:
		break
	case logger.Error:
		logNx.Get().Error(string(p))
	case logger.Warn:
		logNx.Get().Warning(string(p))
	case logger.Info:
		logNx.Get().Info(string(p))
	}
	return len(p), nil
}
