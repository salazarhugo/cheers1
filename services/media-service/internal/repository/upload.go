package repository

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"image"
	"image/jpeg"
)

// UploadImageToCloudStorage uploads data to a Cloud Storage object.
func UploadImageToCloudStorage(
	ctx context.Context,
	bucketName string,
	objectName string,
	img image.Image,
) (*storage.ObjectHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Encode the resized image as JPEG
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, fmt.Errorf("jpeg.Encode: %v", err)
	}

	object := client.Bucket(bucketName).Object(objectName)
	wc := object.NewWriter(ctx)
	defer wc.Close()

	if _, err := wc.Write(buf.Bytes()); err != nil {
		return nil, fmt.Errorf("wc.Write: %v", err)
	}

	return object, nil
}
