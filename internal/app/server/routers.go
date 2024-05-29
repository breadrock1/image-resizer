package server

import (
	"log"

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
// @Failure	503 {object} InternalErrorForm "Server does not available".
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
// @Failure	503 {object} InternalErrorForm "Server does not available".
func (s *Service) TestDownload(c echo.Context) error {
	testImgPath := s.s.GetImagePath("tests.jpg")
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
// @Failure	503 {object} InternalErrorForm "Server does not available".
func (s *Service) Fill(c echo.Context) error {
	reqParams, err := s.extractFormParams(c)
	if err != nil {
		log.Println("failed to extract request: ", err.Error())
		return c.String(400, "Incorrect request params")
	}

	imageAddr := reqParams.ImageAddr.String()
	imageData, storeErr := s.s.DownloadImage(imageAddr, c.Request().Header)
	if storeErr != nil {
		log.Println("failed to download image: ", storeErr.Message)
		return c.String(storeErr.Code, storeErr.Message)
	}

	imageData, err = s.r.ScaleImage(imageData, reqParams.Height, reqParams.Width)
	if err != nil {
		log.Println("failed to scale image: ", err.Error())
		return c.String(400, err.Error())
	}

	imageID, err := s.s.StoreImage(imageData)
	if err != nil {
		log.Println("failed to store scale image: ", err.Error())
	}

	imagePath := s.s.GetImagePath(imageID)
	c.Set("imagePath", imagePath)
	return c.File(imagePath)
}
