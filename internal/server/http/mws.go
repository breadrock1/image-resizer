package http

import (
	"github.com/labstack/echo/v4"
)

func (s *Service) LoadFromCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		imgAddr := c.Request().URL.Path
		if imgPath, exists := s.app.Cacher.GetValue(imgAddr); exists {
			s.app.Logger.Info("loaded from cache image by path: ", imgPath)
			resp := c.Response()
			resp.Status = 304
			return c.File(imgPath)
		}
		return next(c)
	}
}

func (s *Service) StoreToCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		result := next(c)
		if result == nil && c.Response().Status == 200 {
			imagePath := c.Get("imagePath").(string)
			s.app.Cacher.StoreValue(c.Request().URL.Path, imagePath)
			s.app.Logger.Info("stored to cache image by path: ", imagePath)
		}

		return result
	}
}
