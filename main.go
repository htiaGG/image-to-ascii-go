package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

const characters = "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

func main() {
	f, err := os.Open("crescent-moon-png-images-29.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, _, picToAsciiArray, _ := getPixels(f)

	fmt.Printf("%c", picToAsciiArray)
}

func mapTo64Range(value int) int {
	// Perform linear mapping from 0-255 to 0-64
	return int(float64(value) / 255.0 * 64.0)
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
func rgbaToPixelBrightness(r uint32, g uint32, b uint32, a uint32) int {
	brightness := (r/257 + g/257 + b/257) / 3
	return int(brightness)
}

func getPixels(file io.Reader) ([][]Pixel, [][]int, [][]byte, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixelsArray [][]Pixel
	var pixelBrightnessArray [][]int
	var pixelCharsArray [][]byte
	for y := 0; y < height; y++ {
		var row []Pixel
		var rowBrightness []int
		var rowPixelChars []byte
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
			pixelBrightness := rgbaToPixelBrightness(img.At(x, y).RGBA())
			rowBrightness = append(rowBrightness, pixelBrightness)
			rowPixelChars = append(rowPixelChars, characters[mapTo64Range(pixelBrightness)])
		}
		pixelsArray = append(pixelsArray, row)
		pixelBrightnessArray = append(pixelBrightnessArray, rowBrightness)
		pixelCharsArray = append(pixelCharsArray, rowPixelChars)
	}
	return pixelsArray, pixelBrightnessArray, pixelCharsArray, nil
}
