package resizer

import (
	"image-resize-service/internal/config"
	"image-resize-service/internal/resizer/disimage"
)

type Resizer interface {
	ResizeService
}

type ResizeService interface {
	ScaleImage(image []byte, height, width int) ([]byte, error)
}

func New(config *config.ResizerConfig) Resizer {
	resizeService := disimage.New(config)
	return &resizeService
}
