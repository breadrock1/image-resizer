package storage

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"image-resize-service/internal/pkg/config"
)

type Service struct {
	uploadDir string
}

func Create(config *config.StorageConfig) Service {
	return Service{
		uploadDir: config.UploadDirectory,
	}
}

func (s *Service) GetImagePath(imageID string) string {
	return path.Join(s.uploadDir, imageID)
}

func (s *Service) GetImageData(imageID string) ([]byte, error) {
	filePath := path.Join(s.uploadDir, imageID)
	return os.ReadFile(filePath)
}

func (s *Service) StoreImage(image []byte) (string, error) {
	uuidValue, _ := uuid.NewUUID()
	filePath := path.Join(s.uploadDir, uuidValue.String())
	if err := os.WriteFile(filePath, image, 0o600); err != nil {
		return "", err
	}
	return uuidValue.String(), nil
}

func (s *Service) DownloadImage(address string, headers http.Header) ([]byte, *ErrStorage) {
	client := http.Client{}
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", address, nil)
	if err != nil {
		return nil, FromError(err)
	}

	s.setHeaders(req, headers)
	response, err := client.Do(req)
	if err != nil {
		return nil, FromError(err)
	}

	if response.StatusCode > 200 {
		return nil, CreateNew(response.StatusCode, "not founded")
	}

	buffered, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, FromError(err)
	}
	defer func() { _ = response.Body.Close() }()

	return buffered, nil
}

func (s *Service) setHeaders(req *http.Request, headers http.Header) {
	for headKey, headValue := range headers {
		values := strings.Join(headValue, ",")
		req.Header.Add(headKey, values)
	}
}
