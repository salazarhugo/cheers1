package repository

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"image"
	"image/jpeg"
)

// UploadBytesToCloudStorage uploads data to a Cloud Storage object.
func UploadBytesToCloudStorage(
	ctx context.Context,
	bucketName string,
	objectName string,
	bytes []byte,
) (*storage.ObjectHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	object := client.Bucket(bucketName).Object(objectName)
	wc := object.NewWriter(ctx)
	defer wc.Close()

	if _, err := wc.Write(bytes); err != nil {
		return nil, fmt.Errorf("wc.Write: %v", err)
	}

	return object, nil
}

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

	return UploadBytesToCloudStorage(ctx, bucketName, objectName, buf.Bytes())
}
