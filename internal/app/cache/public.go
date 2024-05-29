package cache

import (
	"image-resize-service/internal/pkg/config"
)

type Cache interface {
	Service
}

type Service interface {
	GetValue(address string) (string, bool)
	StoreValue(address string, imagePath string)
}

func New(config *config.CacheConfig) Cache {
	cacheService := Create(config)
	return &cacheService
}
