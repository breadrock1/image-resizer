package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "image-resize-service/docs"
	"image-resize-service/internal/app/cache"
	"image-resize-service/internal/app/resizer"
	"image-resize-service/internal/app/storage"
	"image-resize-service/internal/pkg/config"
)

// @title Swagger Example API
// @version 1.0
// @description Image Resizer API
// @title Image Resizer API
// @host 127.0.0.1:2891
// @BasePath /
// @schemes http

type Service struct {
	address string
	c       cache.Cache
	s       storage.Storage
	r       resizer.Resizer
	e       *echo.Echo
}

func New(config *config.ServerConfig, c cache.Cache, r resizer.Resizer, s storage.Storage) *Service {
	servAddr := fmt.Sprintf("%s:%d", config.Host, config.HostPort)

	servInst := &Service{
		address: servAddr,
		c:       c,
		s:       s,
		r:       r,
		e:       echo.New(),
	}

	servInst.InitEndpoints()
	return servInst
}

func (s *Service) InitEndpoints() {
	s.e.Use(middleware.CORS())
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.GET("/hello/", s.Hello)
	s.e.GET("/tests/tests.jpg", s.TestDownload)
	s.e.GET("/fill/:height/:width/:image", s.Fill, s.LoadFromCache, s.StoreToCache)

	s.e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (s *Service) Start(_ context.Context) error {
	return s.e.Start(s.address)
}

func (s *Service) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
