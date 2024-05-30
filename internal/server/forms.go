package server

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ResponseForm example.
type ResponseForm struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"Done"`
}

// BadRequestForm example.
type BadRequestForm struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"Bad Request message"`
}

// InternalErrorForm example.
type InternalErrorForm struct {
	Status  int    `json:"status" example:"503"`
	Message string `json:"message" example:"Server Error message"`
}

func createStatusResponse(status int, msg string) *ResponseForm {
	return &ResponseForm{Status: status, Message: msg}
}

type FillFormParams struct {
	Height    int
	Width     int
	ImageAddr *url.URL
}

func (s *Service) extractFormParams(c echo.Context) (*FillFormParams, error) {
	var extractErr error
	var width, height int
	var imageAddr *url.URL

	if width, extractErr = strconv.Atoi(c.Param("width")); extractErr != nil {
		return nil, errors.New("incorrect request width param")
	}

	if height, extractErr = strconv.Atoi(c.Param("height")); extractErr != nil {
		return nil, errors.New("incorrect request height param")
	}

	if imageAddr, extractErr = url.Parse(c.Param("image")); extractErr != nil {
		return nil, errors.New("incorrect request image path param")
	}

	return &FillFormParams{
		Height:    height,
		Width:     width,
		ImageAddr: imageAddr,
	}, nil
}
