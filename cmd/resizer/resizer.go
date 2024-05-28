package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"
	"image-resize-service/cmd"
	"image-resize-service/internal/app"
	"image-resize-service/internal/cacher"
	"image-resize-service/internal/logger"
	"image-resize-service/internal/resizer"
	"image-resize-service/internal/server/http"
	"image-resize-service/internal/storage"
)

func main() {
	config := cmd.Execute()

	sCache := cacher.New(&config.Cacher)
	sLog := logger.New(&config.Logger)
	sRes := resizer.New(&config.Resizer)
	sStore := storage.New(&config.Storage)
	sApp := app.New(sCache, sLog, sRes, sStore)

	ctx, cancel := context.WithCancel(context.Background())
	go awaitSystemSignals(cancel)

	server := http.New(&config.Server, sApp)
	go func() {
		if err := server.Start(ctx); err != nil {
			sLog.Error(err.Error())
			cancel()
		}
	}()

	<-ctx.Done()
	cancel()
	shutdownServer(ctx, server)
}

func awaitSystemSignals(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()
}

func shutdownServer(ctx context.Context, server *http.Service) {
	if err := server.Stop(ctx); err != nil {
		log.Warnf("failed to stop server: %s", err)
		return
	}
}
