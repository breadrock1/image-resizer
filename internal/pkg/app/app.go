package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"image-resize-service/internal/app/cache"
	"image-resize-service/internal/app/resizer"
	"image-resize-service/internal/app/server"
	"image-resize-service/internal/app/storage"
	"image-resize-service/internal/pkg/config"
)

type App struct {
	s *server.Service
}

func New(conf *config.Config) *App {
	cacheServ := cache.New(&conf.Cache)
	storeServ := storage.New(&conf.Storage)
	resizeServ := resizer.New(&conf.Resizer)

	httpServer := server.New(&conf.Server, cacheServ, resizeServ, storeServ)

	return &App{s: httpServer}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	go awaitSystemSignals(cancel)

	go func() {
		if err := a.s.Start(ctx); err != nil {
			log.Print(err.Error())
			cancel()
		}
	}()

	<-ctx.Done()
	cancel()
	shutdownServer(ctx, a.s)
}

func awaitSystemSignals(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()
}

func shutdownServer(ctx context.Context, server *server.Service) {
	if err := server.Stop(ctx); err != nil {
		log.Printf("failed to stop server: %s", err)
		return
	}
}
