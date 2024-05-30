package main

import (
	"image-resize-service/cmd"
	"image-resize-service/internal/app"
)

func main() {
	config := cmd.Execute()

	resizeApp := app.New(config)
	resizeApp.Run()
}
