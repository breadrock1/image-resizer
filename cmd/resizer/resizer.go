package main

import (
	"image-resize-service/cmd"
	"image-resize-service/internal/pkg/app"
)

func main() {
	config := cmd.Execute()

	resizeApp := app.New(config)
	resizeApp.Run()
}
