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

// DefaultFilenameTemplate is a mask for generating image file names
const DefaultFilenameTemplate = "image-%d-%d.png"

func main() {
	width := flag.Int("width", 0, "width of the image")
	height := flag.Int("height", 0, "height of the image")
	file := flag.String("file", "", "file to read from")

	flag.Parse()

	hasAnyResolutions := *width != 0 || *height != 0
	hasValidResolutionPair := *width != 0 && *height != 0
	hasFile := *file != ""

	if hasValidResolutionPair == false && hasFile == false {
		log.Fatal("Omitting parameters is not allowed")
	}

	if hasValidResolutionPair && hasFile == false {
		fmt.Println("Creating image from cli provided dimensions")
		createImageFile(width, height)
		return
	}

	if hasFile && hasAnyResolutions {
		log.Fatal("cant specify both file and resolution cli args")
	}

	if hasFile {
		fmt.Println("Processing file input")
		processFile(file)
	}

}

func createImageFile(width *int, height *int) {
	img := createRGBA(*width, *height)
	encodeImage(*width, *height, img)
}

func createRGBA(width int, height int) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2:
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2:
				img.Set(x, y, color.White)
			}
		}
	}
	return img
}

func encodeImage(width int, height int, img *image.RGBA) {
	f, _ := os.Create(fmt.Sprintf(DefaultFilenameTemplate, width, height))
	png.Encode(f, img)
}
