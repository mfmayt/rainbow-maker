// This rainbow maker code is the first code that I felt politic while
// implementing it. This is for people that I love, people of Bogazici,
// LGBTI+ community and all protestors who wants freedom and democrasy
// #ArkadaslarimiziIstiyoruz #BogaziciDirenis #LGBTHaklariInsanHaklari

package main

import (
	"image/color"
	"image/png"
	"os"
)

// Changeable interface allows to change on the image
type Changeable interface {
	Set(x, y int, c color.Color)
}

const (
	leftOffset     int    = 10
	rightOffset    int    = 30
	topOffset      int    = 30
	bottomOffset   int    = 30
	outputFileName string = "rainbow.png"
	inputFileName  string = "logo.png"
)

func main() {
	// Open file
	imgFile, _ := os.Open(inputFileName)
	defer imgFile.Close()

	img, err := png.Decode(imgFile)

	if err != nil {
		panic(err)
	}

	var colors []color.RGBA

	// You can add, remove, update colors from here
	colors = append(colors,
		color.RGBA{255, 0, 24, 255},   // red
		color.RGBA{255, 165, 44, 255}, // orange
		color.RGBA{255, 255, 65, 255}, // yellow
		color.RGBA{0, 128, 24, 255},   // green
		color.RGBA{0, 0, 249, 255},    // blue
		color.RGBA{134, 0, 125, 255})  // purple

	if cimg, ok := img.(Changeable); ok {
		for i := topOffset; i < img.Bounds().Max.Y-bottomOffset; i++ {
			// Select color according its row index
			colorOrder := (i * len(colors)) / img.Bounds().Max.Y

			for j := leftOffset; j < img.Bounds().Max.X-rightOffset; j++ {
				r, g, b, a := img.At(j, i).RGBA()

				// Check is pixel transparent empty
				if r != 0 || g != 0 || b != 0 || a != 0 {
					cimg.Set(j, i, colors[colorOrder])
				}
			}
		}

		// Write output to file
		outFile, err := os.Create(outputFileName)
		defer outFile.Close()

		if err != nil {
			panic(err)
		}

		png.Encode(outFile, img)
	}
}
