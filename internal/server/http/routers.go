package http

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Hello
// @Summary Hello
// @Router /hello/ [get]
// @Tags hello
// @Description Check service is available
// @ID hello
// @Produce  json
// @Success 200 {object} ResponseForm "Ok"
// @Failure	503 {object} ServerErrorForm "Server does not available".
func (s *Service) Hello(c echo.Context) error {
	okResp := createStatusResponse(200, "Ok")
	return c.JSON(200, okResp)
}

// TestDownload
// @Summary TestDownload
// @Router /tests/tests.jpg [get]
// @Tags tests
// @Description Resize image by URL address
// @ID tests-download
// @Produce  multipart/form
// @Success 200 {object} ResponseForm "Ok"
// @Success 400 {object} BadRequestForm "Client error"
// @Success 502 {object} BadRequestForm "Gateway error"
// @Failure	503 {object} ServerErrorForm "Server does not available".
func (s *Service) TestDownload(c echo.Context) error {
	testImgPath := s.app.Storage.GetImagePath("tests.jpg")
	return c.File(testImgPath)
}

// Fill
// @Summary Fill
// @Router /fill/{height}/{width}/{image} [get]
// @Tags fill
// @Description Resize image by URL address
// @ID fill
// @Accept  multipart/form
// @Param height path string true "Height"
// @Param width path string true "Width"
// @Param image path string true "Image URL address"
// @Success 200 {object} ResponseForm "Ok"
// @Success 400 {object} BadRequestForm "Client error"
// @Success 502 {object} BadRequestForm "Gateway error"
// @Failure	503 {object} ServerErrorForm "Server does not available".
func (s *Service) Fill(c echo.Context) error {
	reqParams, err := s.extractFormParams(c)
	if err != nil {
		s.app.Logger.Error("failed to extract request: ", err.Error())
		return c.String(400, "Incorrect request params")
	}

	imageAddr := reqParams.ImageAddr.String()
	imageData, storeErr := s.app.Storage.DownloadImage(imageAddr, c.Request().Header)
	if storeErr != nil {
		s.app.Logger.Error("failed to download image: ", storeErr.Message)
		return c.String(storeErr.Code, storeErr.Message)
	}

	imageData, err = s.app.Resizer.ScaleImage(imageData, reqParams.Height, reqParams.Width)
	if err != nil {
		s.app.Logger.Error("failed to scale image: ", err.Error())
		return c.String(400, err.Error())
	}

	imageID, err := s.app.Storage.StoreImage(imageData)
	if err != nil {
		s.app.Logger.Error("failed to store scale image: ", err.Error())
	}

	imagePath := s.app.Storage.GetImagePath(imageID)
	c.Set("imagePath", imagePath)
	return c.File(imagePath)
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
