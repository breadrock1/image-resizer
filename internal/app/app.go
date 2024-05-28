package app

import (
	"image-resize-service/internal/cacher"
	"image-resize-service/internal/logger"
	"image-resize-service/internal/resizer"
	"image-resize-service/internal/storage"
)

type App struct {
	Cacher  cacher.Cacher
	Logger  logger.Logger
	Storage storage.Storage
	Resizer resizer.Resizer
}

func New(sCache cacher.Cacher, sLog logger.Logger, sRes resizer.Resizer, sStore storage.Storage) *App {
	return &App{
		Cacher:  sCache,
		Logger:  sLog,
		Resizer: sRes,
		Storage: sStore,
	}
}
