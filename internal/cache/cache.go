package cache

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"image-resize-service/internal/config"
)

type Memcache struct {
	expire      time.Duration
	cacheClient *expirable.LRU[string, string]
}

func Create(config *config.CacheConfig) Memcache {
	expireTime := time.Duration(config.ExpireSeconds) * time.Second
	cacheInst := expirable.NewLRU[string, string](config.CapacityValues, nil, expireTime)
	return Memcache{
		expire:      expireTime,
		cacheClient: cacheInst,
	}
}

func (m *Memcache) GetValue(address string) (string, bool) {
	value, ok := m.cacheClient.Get(address)
	if !ok {
		return "", false
	}
	return value, true
}

func (m *Memcache) StoreValue(address string, imagePath string) {
	m.cacheClient.Add(address, imagePath)
}
