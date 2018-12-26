package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func toHex(r, g, b, a uint32) string {
	return fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: colorcount file")
		fmt.Println("Counts colors in an image [jpg, png, gif].")
		os.Exit(1)
	}
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	colors := make(map[string]bool)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colors[toHex(img.At(x, y).RGBA())] = true
		}
	}
	fmt.Println(len(colors))
	os.Exit(0)
}
