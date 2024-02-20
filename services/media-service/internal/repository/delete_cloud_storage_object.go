package repository

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
)

func DeleteCloudStorageObject(
	bucketName string,
	path string,
) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Get image from cloud storage
	err = client.Bucket(bucketName).Object(path).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
