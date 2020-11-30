package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	width := flag.Int("width", 0, "width of the image")
	height := flag.Int("height", 0, "height of the image")

	flag.Parse()

	if *width == 0 || *height == 0 {
		log.Fatal("Zero values are not allowed")
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{*width, *height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < *width; x++ {
		for y := 0; y < *height; y++ {
			switch {
			case x < *width/2 && y < *height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= *width/2 && y >= *height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(fmt.Sprintf("image-%d-%d.png", *width, *height))
	png.Encode(f, img)
}
