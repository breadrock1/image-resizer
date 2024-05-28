package service

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	address string
	server  *echo.Echo
}

func New(servPort int) *Service {
	servAddr := fmt.Sprintf("localhost:%d", servPort)
	servInst := &Service{
		address: servAddr,
		server:  echo.New(),
	}

	servInst.InitEndpoints()
	return servInst
}

func (s *Service) InitEndpoints() {
	s.server.Use(middleware.CORS())

	s.server.GET("/image/image", s.ReturnImage)
	s.server.GET("/image/not_image", s.ReturnNonImage)
	s.server.GET("/image/internal_error", s.InternalServerError)
}

func (s *Service) Start(_ context.Context) error {
	return s.server.Start(s.address)
}

func (s *Service) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
