package logger

import (
	"image-resize-service/internal/config"
	"image-resize-service/internal/logger/local"
)

type Logger interface {
	ServiceTrait
}

type ServiceTrait interface {
	Info(msg ...string)
	Warn(msg ...string)
	Error(msg ...string)
}

func New(config *config.LoggerConfig) Logger {
	logService := local.New(config)
	return &logService
}
