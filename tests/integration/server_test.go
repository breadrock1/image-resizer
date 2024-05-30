package integration

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"image-resize-service/internal/cache"
	"image-resize-service/internal/config"
	"image-resize-service/internal/resizer"
	"image-resize-service/internal/server"
	"image-resize-service/internal/storage"
	"image-resize-service/tests/integration/service"
)

const (
	MockServerPort    = 7453
	MockServerAddr    = "http://localhost:7453/image"
	ResizerServerAddr = "http://localhost:2891/fill"
)

func TestBaseCases(t *testing.T) {
	go LaunchResizerService()
	go LaunchMockService()

	t.Run("Check cache processing", func(t *testing.T) {
		addr := fmt.Sprintf("%s/64/64/%s/image", ResizerServerAddr, MockServerAddr)

		startTime := time.Now()
		resp, err := SendRequest(addr)
		stopTime := time.Now()
		require.NoError(t, err, "unexpected error")
		require.NotEmpty(t, resp.Body, "does not been empty response")
		_ = resp.Body.Close()

		procDuration := stopTime.Nanosecond() - startTime.Nanosecond()

		startTime = time.Now()
		resp, err = SendRequest(addr)
		stopTime = time.Now()
		require.NoError(t, err, "unexpected error")
		require.NotEmpty(t, resp.Body, "does not been empty response")
		_ = resp.Body.Close()

		cacheDuration := stopTime.Nanosecond() - startTime.Nanosecond()

		require.Equal(t, true, cacheDuration < procDuration)
	})

	t.Run("Remote server does not exist", func(t *testing.T) {
		fakeServiceAddr := "http://localhost:1111"
		addr := fmt.Sprintf("%s/64/64/%s/image", ResizerServerAddr, fakeServiceAddr)

		resp, _ := SendRequest(addr)
		require.Equal(t, 502, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("Returned 404 error from remote server", func(t *testing.T) {
		addr := fmt.Sprintf("%s/64/64/%s/not_founded", ResizerServerAddr, MockServerAddr)

		resp, _ := SendRequest(addr)
		require.Equal(t, 404, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("Returned not image from remote server", func(t *testing.T) {
		addr := fmt.Sprintf("%s/64/64/%s/not_image", ResizerServerAddr, MockServerAddr)

		resp, _ := SendRequest(addr)
		require.Equal(t, 400, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("Returned server error", func(t *testing.T) {
		addr := fmt.Sprintf("%s/64/64/%s/internal_error", ResizerServerAddr, MockServerAddr)

		resp, _ := SendRequest(addr)
		require.Equal(t, 500, resp.StatusCode)
		_ = resp.Body.Close()
	})
}

func LaunchMockService() {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	mockServer := service.New(MockServerPort)
	go func() {
		if err := mockServer.Start(ctx); err != nil {
			log.Println(err.Error())
			cancel()
		}
	}()

	<-ctx.Done()
	cancel()
	_ = mockServer.Stop(ctx)
}

func LaunchResizerService() {
	appConfig, err := config.NewConfig("../resources/config.toml")
	if err != nil {
		log.Fatalf("Failed while parsing config file: %s", err)
	}

	sCache := cache.New(&appConfig.Cache)
	sRes := resizer.New(&appConfig.Resizer)
	sStore := storage.New(&appConfig.Storage)

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	httpServer := server.New(&appConfig.Server, sCache, sRes, sStore)
	go func() {
		if err := httpServer.Start(ctx); err != nil {
			log.Println(err.Error())
			cancel()
		}
	}()

	<-ctx.Done()
	cancel()
	_ = httpServer.Stop(ctx)
}

func SendRequest(address string) (*http.Response, error) {
	client := http.Client{}
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", address, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
