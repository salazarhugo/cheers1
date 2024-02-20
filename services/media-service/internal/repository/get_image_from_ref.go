package repository

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"image"
)

func GetCloudStorageImage(bucketName string, path string) (*image.Image, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Get image from cloud storage
	rc, err := client.Bucket(bucketName).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	// Decode the image
	img, _, err := image.Decode(rc)
	if err != nil {
		return nil, fmt.Errorf("image.Decode: %v", err)
	}

	return &img, nil
}
