package resizer

import (
	"image-resize-service/internal/config"
)

type Resizer interface {
	ResizeService
}

type ResizeService interface {
	ScaleImage(image []byte, height, width int) ([]byte, error)
}

func New(config *config.ResizerConfig) Resizer {
	resizeService := CreateResizer(config)
	return &resizeService
}
