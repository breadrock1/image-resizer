package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"image-resize-service/internal/config"
	"image-resize-service/internal/storage"
)

const UploadsDir = "../../../uploads"
const CorrectURL = "https://images.rawpixel.com/image_png_800/test.png"
const ImgURL = "https://img.freepik.com/free-psd/mix-fruits-png-isolated-transparent-background_191095-9865.jpg"
const NonImgURL = "https://images.rawl.com/image_png_800/nBuZw.png"

func TestBaseStorage(t *testing.T) {
	storeConfig := config.StorageConfig{UseFilesystem: true, UploadDirectory: UploadsDir}
	storeService := storage.New(&storeConfig)

	t.Run("Extract URL", func(t *testing.T) {
		_, err := storeService.ExtractImageURL(CorrectURL)
		assert.NoError(t, err, "unexpected error")
	})

	t.Run("Extract incorrect URL", func(t *testing.T) {
		_, err := storeService.ExtractImageURL("")
		assert.Error(t, err, "expected error")
	})

	t.Run("Get image", func(t *testing.T) {
		_, err := storeService.GetImage("test.jpg")
		assert.NoError(t, err, "unexpected error")
	})

	t.Run("Get non existing image", func(t *testing.T) {
		_, err := storeService.GetImage("some.jpg")
		assert.Error(t, err, "expected error")
	})

	t.Run("Store image", func(t *testing.T) {
		imgData, err := storeService.GetImage("test.jpg")
		assert.NoError(t, err, "unexpected error")

		imgUUID, err := storeService.StoreImage(imgData)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, len(imgUUID) > 0, true)
	})
}

func TestIncorrectUploadDir(t *testing.T) {
	storeConfig := config.StorageConfig{UseFilesystem: true, UploadDirectory: "./nonexist"}
	storeService := storage.New(&storeConfig)

	imgPath := fmt.Sprintf("%s/%s", UploadsDir, "test.png")
	testImgData, _ := storeService.GetImage(imgPath)

	t.Run("Get non existing image", func(t *testing.T) {
		_, err := storeService.GetImage("test.jpg")
		assert.Error(t, err, "expected error")
	})

	t.Run("Store image to non existing dir", func(t *testing.T) {
		imgUUID, err := storeService.StoreImage(testImgData)
		assert.Error(t, err, "expected error")
		assert.Equal(t, len(imgUUID) < 1, true)
	})
}

func TestDownload(t *testing.T) {
	storeConfig := config.StorageConfig{UseFilesystem: true, UploadDirectory: UploadsDir}
	storeService := storage.New(&storeConfig)

	t.Run("Download image", func(t *testing.T) {
		_, err := storeService.DownloadImage(ImgURL, make(map[string][]string))
		assert.Empty(t, err, "unexpected error")
	})

	t.Run("Download non existing image", func(t *testing.T) {
		_, err := storeService.DownloadImage(NonImgURL, make(map[string][]string))
		assert.Empty(t, err, "expected error")
	})
}
