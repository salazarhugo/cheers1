package repository

import (
	"github.com/disintegration/imaging"
	"image"
)

// SquareResizeImage resizes the image to the specified width and height.
func SquareResizeImage(
	img image.Image,
) (image.Image, error) {
	// Get the dimensions of the image
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// Calculate the size for center cropping
	var size int
	if width > height {
		size = height
	} else {
		size = width
	}

	// Center crop the image
	croppedImg := imaging.CropCenter(img, size, size)

	// Resize the cropped image to 1080x1080
	resizedImg := imaging.Resize(croppedImg, 1080, 1080, imaging.Lanczos)

	return resizedImg, nil
}
