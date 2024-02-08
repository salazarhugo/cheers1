package repository

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
)

// UploadToCloudStorage uploads data to a Cloud Storage object.
func UploadToCloudStorage(
	ctx context.Context,
	bucketName,
	objectName string,
	data []byte,
) (*storage.ObjectHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	object := client.Bucket(bucketName).Object(objectName)
	wc := object.NewWriter(ctx)
	defer wc.Close()

	if _, err := wc.Write(data); err != nil {
		return nil, fmt.Errorf("wc.Write: %v", err)
	}

	return object, nil
}
