package service

import (
	"errors"
	"path"

	"github.com/labstack/echo/v4"
)

const ResourcesDir = "../resources"

func (s *Service) InternalServerError(_ echo.Context) error {
	errResp := errors.New("something gone wrong")
	return errResp
}

func (s *Service) ReturnNonImage(c echo.Context) error {
	filePath := path.Join(ResourcesDir, "test.txt")
	return c.File(filePath)
}

func (s *Service) ReturnImage(c echo.Context) error {
	filePath := path.Join(ResourcesDir, "test.jpg")
	return c.File(filePath)
}
