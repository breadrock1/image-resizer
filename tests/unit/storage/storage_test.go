package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"image-resize-service/internal/app/storage"
	"image-resize-service/internal/pkg/config"
)

const (
	uploadDirectory     = "../../../uploads"
	existingImageURL    = "https://upload.wikimedia.org/wikipedia/commons/3/3f/JPEG_example_flower.jpg"
	nonExistingImageURL = "https://images.rawl.com/image_png_800/nBuZw.png"
)

func TestBaseStorage(t *testing.T) {
	storeConfig := config.StorageConfig{UseFilesystem: true, UploadDirectory: uploadDirectory}
	storeService := storage.New(&storeConfig)

	t.Run("Get image", func(t *testing.T) {
		_, err := storeService.GetImageData("test.jpg")
		assert.NoError(t, err, "unexpected error")
	})

	t.Run("Get non existing image", func(t *testing.T) {
		_, err := storeService.GetImageData("some.jpg")
		assert.Error(t, err, "expected error")
	})

	t.Run("Store image", func(t *testing.T) {
		imgData, err := storeService.GetImageData("test.jpg")
		assert.NoError(t, err, "unexpected error")

		imgUUID, err := storeService.StoreImage(imgData)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, len(imgUUID) > 0, true)
	})
}

func TestIncorrectUploadDir(t *testing.T) {
	storeConfig := config.StorageConfig{UseFilesystem: true, UploadDirectory: "./nonexist"}
	storeService := storage.New(&storeConfig)

	imgPath := fmt.Sprintf("%s/%s", uploadDirectory, "test.png")
	testImgData, _ := storeService.GetImageData(imgPath)

	t.Run("Get non existing image", func(t *testing.T) {
		_, err := storeService.GetImageData("test.jpg")
		assert.Error(t, err, "expected error")
	})

	t.Run("Store image to non existing dir", func(t *testing.T) {
		imgUUID, err := storeService.StoreImage(testImgData)
		assert.Error(t, err, "expected error")
		assert.Equal(t, len(imgUUID) < 1, true)
	})
}

func TestDownload(t *testing.T) {
	storeConfig := config.StorageConfig{UseFilesystem: true, UploadDirectory: uploadDirectory}
	storeService := storage.New(&storeConfig)

	t.Run("Download image", func(t *testing.T) {
		_, err := storeService.DownloadImage(existingImageURL, make(map[string][]string))
		assert.Empty(t, err, "unexpected error")
	})

	t.Run("Download non existing image", func(t *testing.T) {
		_, err := storeService.DownloadImage(nonExistingImageURL, make(map[string][]string))
		assert.Equal(t, 502, err.Code, "expected error")
	})
}
