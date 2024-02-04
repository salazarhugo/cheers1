package repository

import (
	"github.com/nfnt/resize"
	"image"
)

// ResizeImage resizes the image to the specified width and height.
func ResizeImage(img image.Image, width, height uint) (image.Image, error) {
	return resize.Resize(width, height, img, resize.Lanczos3), nil
}
