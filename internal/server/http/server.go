package http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "image-resize-service/docs"
	"image-resize-service/internal/app"
	"image-resize-service/internal/config"
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
	app     *app.App
	server  *echo.Echo
}

func New(config *config.ServerConfig, sApp *app.App) *Service {
	servAddr := fmt.Sprintf("%s:%d", config.Host, config.HostPort)

	servInst := &Service{
		address: servAddr,
		app:     sApp,
		server:  echo.New(),
	}

	servInst.InitEndpoints()
	return servInst
}

func (s *Service) InitEndpoints() {
	s.server.Use(middleware.CORS())
	s.server.Use(middleware.Logger())

	s.server.GET("/hello/", s.Hello)
	s.server.GET("/tests/tests.jpg", s.TestDownload)
	s.server.GET("/fill/:height/:width/:image", s.Fill, s.LoadFromCache, s.StoreToCache)

	s.server.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (s *Service) Start(_ context.Context) error {
	return s.server.Start(s.address)
}

func (s *Service) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
