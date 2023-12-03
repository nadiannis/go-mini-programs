package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Image path is required")
	}

	filePath := os.Args[1]
	img := loadImage(filePath)

	imgSize := img.Bounds().Size()
	newImg := createImage(imgSize)

	convertToBlackAndWhite(img, newImg)
	saveImage(newImg, filePath)
}

func createImage(size image.Point) *image.RGBA {
	rect := image.Rect(0, 0, size.X, size.Y)
	return image.NewRGBA(rect)
}

func loadImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if !isValidExtension(path) {
		log.Fatalln("Invalid format. You can only convert JPG/ JPEG/ PNG.")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	return img
}

func saveImage(img *image.RGBA, path string) {
	ext := filepath.Ext(path)
	name := strings.TrimSuffix(filepath.Base(path), ext)
	newPath := fmt.Sprintf("%s/%s_bw%s", filepath.Dir(path), name, ext)
	
	file, err := os.Create(newPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if ext == ".jpg" || ext == ".jpeg" {
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			log.Fatalln(err)
		}
	} else if ext == ".png" {
		err = png.Encode(file, img)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func convertToBlackAndWhite(img image.Image, newImg *image.RGBA) {
	size := img.Bounds().Size()

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixelRGB := img.At(x, y)
			pixelRGBA := color.RGBAModel.Convert(pixelRGB)
			originalColor, ok := pixelRGBA.(color.RGBA)
			if !ok {
				log.Fatalln("Pixel is not in RGBA")
			}

			r := float64(originalColor.R) * 0.421
			g := float64(originalColor.G) * 0.405 
			b := float64(originalColor.B) * 0.467

			gray := uint8(r + g + b / 3)
			newColor := color.RGBA{
				R: gray,
				G: gray,
				B: gray,
				A: originalColor.A,
			}

			newImg.Set(x, y, newColor)
		}
	}
}

func isValidExtension(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
		return true
	}
	return false
}