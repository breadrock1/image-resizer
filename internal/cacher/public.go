package cacher

import (
	"image-resize-service/internal/cacher/memcache"
	"image-resize-service/internal/config"
)

type Cacher interface {
	CacheService
}

type CacheService interface {
	GetValue(address string) (string, bool)
	StoreValue(address string, imagePath string)
}

func New(config *config.CacheConfig) Cacher {
	cacheService := memcache.New(config)
	return &cacheService
}
