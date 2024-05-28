package fs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"image-resize-service/internal/config"
	"image-resize-service/internal/storage/storage_err"
)

type Service struct {
	uploadDir string
}

func New(config *config.StorageConfig) Service {
	return Service{
		uploadDir: config.UploadDirectory,
	}
}

func (s *Service) GetImagePath(imageID string) string {
	return path.Join(s.uploadDir, imageID)
}

func (s *Service) StoreImage(image []byte) (string, error) {
	uuidValue, _ := uuid.NewUUID()
	filePath := path.Join(s.uploadDir, uuidValue.String())
	if err := os.WriteFile(filePath, image, 0o600); err != nil {
		return "", err
	}
	return uuidValue.String(), nil
}

func (s *Service) GetImage(imageID string) ([]byte, error) {
	filePath := path.Join(s.uploadDir, imageID)
	return os.ReadFile(filePath)
}

func (s *Service) ExtractImageURL(address string) (string, error) {
	if len(address) < 1 {
		return "", errors.New("empty image URL address")
	}

	imageURL, err := url.Parse(address)
	if err != nil {
		return "", err
	}

	return imageURL.Path, nil
}

func (s *Service) DownloadImage(address string, headers http.Header) ([]byte, *storageerr.ErrStorage) {
	client := http.Client{}
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", address, nil)
	if err != nil {
		return nil, storageerr.FromError(err)
	}

	s.setHeaders(req, headers)
	response, err := client.Do(req)
	if err != nil {
		return nil, storageerr.FromError(err)
	}

	if response.StatusCode > 200 {
		return nil, storageerr.New(response.StatusCode, "not founded")
	}

	buffered, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, storageerr.FromError(err)
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
