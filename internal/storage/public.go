package storage

import (
	"net/http"

	"image-resize-service/internal/config"
)

type Storage interface {
	StoreService
	LoadService
}

type StoreService interface {
	GetImagePath(imageID string) string
	StoreImage(image []byte) (string, error)
	GetImageData(imageID string) ([]byte, error)
}

type LoadService interface {
	DownloadImage(imgAddr string, headers http.Header) ([]byte, *ErrStorage)
}

func New(config *config.StorageConfig) Storage {
	storeService := Create(config)
	return &storeService
}
