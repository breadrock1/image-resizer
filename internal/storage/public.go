package storage

import (
	"net/http"

	"image-resize-service/internal/config"
	"image-resize-service/internal/storage/fs"
	"image-resize-service/internal/storage/storage_err"
)

type Storage interface {
	StoreService
	LoadService
}

type StoreService interface {
	GetImagePath(imageID string) string
	StoreImage(image []byte) (string, error)
	GetImage(imageID string) ([]byte, error)
}

type LoadService interface {
	ExtractImageURL(address string) (string, error)
	DownloadImage(imgAddr string, headers http.Header) ([]byte, *storageerr.ErrStorage)
}

func New(config *config.StorageConfig) Storage {
	storeService := fs.New(config)
	return &storeService
}
