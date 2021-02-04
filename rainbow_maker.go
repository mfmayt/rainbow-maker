// This rainbow maker code is the first code that I felt politic while
// implementing it. This is for people that I love, people of Bogazici,
// LGBTI+ community and all protestors who wants freedom and democrasy
// #ArkadaslarimiziIstiyoruz #BogaziciDirenis #LGBTHaklariInsanHaklari

package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

// Changeable interface allows to change on the image
type Changeable interface {
	Set(x, y int, c color.Color)
}

const (
	outputFileName string = "rainbow.png"
	inputFileName  string = "logo.png"
)

// You can add, remove, update colors from here
var colors = []color.RGBA{
	{255, 0, 24, 255},   // red
	{255, 165, 44, 255}, // orange
	{255, 255, 65, 255}, // yellow
	{0, 128, 24, 255},   // green
	{0, 0, 249, 255},    // blue
	{134, 0, 125, 255}}  // purple

func main() {
	// Open file
	imgFile, _ := os.Open(inputFileName)
	defer imgFile.Close()

	img, err := png.Decode(imgFile)

	if err != nil {
		panic(err)
	}

	str := `Type '1' or '2':
	'1' if you want to change colors of only non-transparent pixels, 
	'2' if you want to change colors of only transparent pixels`
	fmt.Printf(str + "\n")
	var choice int
	fmt.Scanf("%d", &choice)

	if cimg, ok := img.(Changeable); ok {
		for i := 0; i < img.Bounds().Max.Y; i++ {
			// Select color according its row index
			colorOrder := (i * len(colors)) / img.Bounds().Max.Y

			for j := 0; j < img.Bounds().Max.X; j++ {
				r, g, b, a := img.At(j, i).RGBA()

				switch choice {
				case 1:
					// Check is pixel transparent
					if r != 0 || g != 0 || b != 0 || a != 0 {
						cimg.Set(j, i, colors[colorOrder])
					}
					break
				case 2:
					// Check is pixel not transparent or empty
					// TODO: noise cancel algorithm can be used here
					if !(r == 0 && g == 0 && b == 0) || a == 0 {
						cimg.Set(j, i, colors[colorOrder])
					}
					break
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
