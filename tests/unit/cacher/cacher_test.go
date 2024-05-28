package cacher

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"image-resize-service/internal/cacher"
	"image-resize-service/internal/config"
)

const CapacityValues = 4
const ExpireSeconds = 3
const URLAddress = "http://localhost:2891/15/15"

func TestBaseCacher(t *testing.T) {
	cacherConfig := config.CacheConfig{ExpireSeconds: ExpireSeconds, CapacityValues: CapacityValues}
	cacherService := cacher.New(&cacherConfig)

	t.Run("Base functionality", func(t *testing.T) {
		for i := 0; i <= CapacityValues; i++ {
			imgName, urlAddr := buildImageName(i)
			cacherService.StoreValue(urlAddr, imgName)
		}

		time.Sleep(2 * time.Second)

		_, firstURLAddr := buildImageName(1)
		_, exists := cacherService.GetValue(firstURLAddr)
		assert.Equal(t, exists, true)
	})

	t.Run("Expire elements", func(t *testing.T) {
		for i := 0; i <= CapacityValues; i++ {
			imgName, urlAddr := buildImageName(i)
			cacherService.StoreValue(urlAddr, imgName)
		}

		time.Sleep(2 * time.Second)

		_, firstURLAddr := buildImageName(1)
		_, exists := cacherService.GetValue(firstURLAddr)
		assert.Equal(t, exists, true)

		time.Sleep(4 * time.Second)

		_, exists = cacherService.GetValue(firstURLAddr)
		assert.Equal(t, exists, false)
	})
}

func buildImageName(index int) (string, string) {
	imgName := fmt.Sprintf("test_%d.jpg", index)
	urlAddr := fmt.Sprintf("%s/%s", URLAddress, imgName)
	return imgName, urlAddr
}
