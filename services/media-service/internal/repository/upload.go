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
) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	defer wc.Close()

	if _, err := wc.Write(data); err != nil {
		return fmt.Errorf("wc.Write: %v", err)
	}

	return nil
}
