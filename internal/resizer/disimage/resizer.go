package disimage

import (
	"bytes"
	"errors"
	"image"

	"github.com/disintegration/imaging"
	"image-resize-service/internal/config"
)

type Service struct {
	TargetQuality int
}

func New(config *config.ResizerConfig) Service {
	return Service{
		TargetQuality: config.TargetQuality,
	}
}

func (s *Service) ScaleImage(image []byte, height, width int) ([]byte, error) {
	r := bytes.NewReader(image)
	decodedImage, err := imaging.Decode(r, imaging.AutoOrientation(true))
	if err != nil {
		return nil, err
	}

	if err = s.checkImageSize(decodedImage, width, height); err != nil {
		return nil, err
	}

	resized := imaging.Resize(decodedImage, width, height, imaging.MitchellNetravali)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, resized, imaging.JPEG, imaging.JPEGQuality(s.TargetQuality))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *Service) checkImageSize(image image.Image, height, width int) error {
	imgWidth := image.Bounds().Size().X
	imgHeight := image.Bounds().Size().Y
	if imgWidth < width || imgHeight < height {
		return errors.New("incorrect passed width and height to scale")
	}

	return nil
}
