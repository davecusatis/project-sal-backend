package image

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"mime/multipart"
	"os"
)

const (
	numberOfSubImages = 9
	subImgWidth       = 128
	subImgHeight      = 128
	newImageWidth     = 128
	newImageHeight    = 128 * numberOfSubImages
)

var validImageFormats = []string{"png", "jpeg", "jpg"}

func validateFormat(format string) bool {
	found := false
	for _, f := range validImageFormats {
		if f == format {
			found = true
			break
		}
	}
	return found
}

func ValidateImageFromFile(f *multipart.FileHeader) bool {
	reader, err := f.Open()
	if err != nil {
		log.Printf("Error opening file")
		return false
	}
	defer reader.Close()

	config, format, err := image.Decode(reader)
	if err != nil {
		log.Printf("Error decoding %s file: %s", f.Filename, err)
		return false
	}
	x := config.Bounds().Dx()
	y := config.Bounds().Dy()
	return (x <= 128 && x > 0) && (y <= 128 && y > 0) && validateFormat(format)
}

func GenerateImageFromURLS() {
	// create new image to be written
	out := image.NewRGBA(image.Rect(0, 0, newImageWidth, newImageHeight))

	subImages := []string{
		"img/iconBar.png",
		"img/iconBell.png",
		"img/iconCherries.png",
		"img/iconCoin.png",
		"img/iconDiamond.png",
		"img/iconHorseshoe.png",
		"img/iconLime.png",
		"img/iconPlum.png",
		"img/iconSeven.png",
	}

	// for each image
	// 		download then load image
	// 		copy image onto new image
	newY := 0
	for _, img := range subImages {
		in, err := os.Open(img)
		if err != nil {
			log.Fatalf("Fatal error opening file to load: %s", err)
		}
		defer in.Close()

		src, _, err := image.Decode(in)
		if err != nil {
			log.Fatalf("Fatal error decoding image file: %s", err)
		}

		for x := 0; x < subImgWidth; x++ {
			for y := 0; y < subImgHeight; y++ {
				toCopy := src.At(x, y)
				out.Set(x, y+newY, toCopy)
			}
		}
		newY = newY + subImgHeight
	}

	// save new image
	f, err := os.OpenFile("img/out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
	}
	defer f.Close()
	png.Encode(f, out)

	// upload to s3
	// clean up
}
