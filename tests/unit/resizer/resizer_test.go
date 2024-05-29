package resizer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"image-resize-service/internal/app/resizer"
	"image-resize-service/internal/pkg/config"
)

func TestResizer(t *testing.T) {
	resizerConfig := config.ResizerConfig{TargetQuality: 90}
	resizerService := resizer.New(&resizerConfig)

	t.Run("Base functionality", func(t *testing.T) {
		imgPath := "../../../uploads/test.jpg"
		imgData, err := os.ReadFile(imgPath)
		assert.NoError(t, err, "failed to read testcase img file")

		_, err = resizerService.ScaleImage(imgData, 64, 64)
		assert.NoError(t, err, "failed to scale testcase img file")
	})

	t.Run("Non JPEG img format", func(t *testing.T) {
		imgPath := "../../../uploads/butterfly.webp"
		imgData, err := os.ReadFile(imgPath)
		assert.NoError(t, err, "failed to read testcase img file")

		_, err = resizerService.ScaleImage(imgData, 64, 64)
		assert.Error(t, err, "this must be a error")
	})
}
