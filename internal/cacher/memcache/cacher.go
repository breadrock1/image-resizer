package memcache

import (
	"fmt"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"image-resize-service/internal/config"
)

type Service struct {
	expire time.Duration
	cacher *expirable.LRU[string, string]
}

func New(config *config.CacheConfig) Service {
	expireTime := time.Duration(config.ExpireSeconds) * time.Second
	cacheInst := expirable.NewLRU[string, string](config.CapacityValues, nil, expireTime)
	return Service{
		expire: expireTime,
		cacher: cacheInst,
	}
}

func (s *Service) GetValue(address string) (string, bool) {
	value, ok := s.cacher.Get(address)
	if !ok {
		return "", false
	}
	return value, true
}

func (s *Service) StoreValue(address string, imagePath string) {
	fmt.Println(len(s.cacher.Keys()))
	s.cacher.Add(address, imagePath)
}
