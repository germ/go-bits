// Package pngGen generates a random png image of the given dimensions
// using only the CPU for processing
package pngGen

import (
	"germ/puuid"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Generate creates a png with data according to draw
func Generate(width, height int) (string, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//Spawning a thread for every pixel would be ideal but it uses too many threads
	//A comprimise is to spawn a thread for every line
	for y := 0; y < height; y++ {
		go draw(img, y)
	}

	f, err := os.Create("images/" + puuid.Generate() + ".png")

	if err != nil {
		return "", err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

// draw contains the logic for drawing the png
func draw(img *image.RGBA, y int) {
	for i := 0; i < img.Rect.Dx(); i++ {
		var r, g, b uint8
		switch (i/10 + y/10) % 3 {
		case 0:
			r, g, b = uint8(255), uint8(0), uint8(0)
		case 1:
			r, g, b = uint8(0), uint8(255), uint8(0)
		case 2:
			r, g, b = uint8(0), uint8(0), uint8(255)
		}

		img.Set(i, y, color.RGBA{r, g, b, uint8(255)})
	}
}
