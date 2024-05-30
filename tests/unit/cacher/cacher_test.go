package cacher

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"image-resize-service/internal/cache"
	"image-resize-service/internal/config"
)

const (
	capacityValues = 4
	expireSeconds  = 3
	addressURL     = "http://localhost:2891/15/15"
)

func TestBaseCache(t *testing.T) {
	cacheConfig := config.CacheConfig{ExpireSeconds: expireSeconds, CapacityValues: capacityValues}
	cacheService := cache.New(&cacheConfig)

	t.Run("Base functionality", func(t *testing.T) {
		for i := 0; i <= capacityValues; i++ {
			imgName, urlAddr := buildImageName(i)
			cacheService.StoreValue(urlAddr, imgName)
		}

		time.Sleep(2 * time.Second)

		_, firstURLAddr := buildImageName(1)
		_, exists := cacheService.GetValue(firstURLAddr)
		assert.Equal(t, exists, true)
	})

	t.Run("Expire elements", func(t *testing.T) {
		for i := 0; i <= capacityValues; i++ {
			imgName, urlAddr := buildImageName(i)
			cacheService.StoreValue(urlAddr, imgName)
		}

		time.Sleep(2 * time.Second)

		_, firstURLAddr := buildImageName(1)
		_, exists := cacheService.GetValue(firstURLAddr)
		assert.Equal(t, exists, true)

		time.Sleep(4 * time.Second)

		_, exists = cacheService.GetValue(firstURLAddr)
		assert.Equal(t, exists, false)
	})
}

func buildImageName(index int) (string, string) {
	imgName := fmt.Sprintf("test_%d.jpg", index)
	urlAddr := fmt.Sprintf("%s/%s", addressURL, imgName)
	return imgName, urlAddr
}
