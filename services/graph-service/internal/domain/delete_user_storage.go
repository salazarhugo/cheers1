package domain

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
)

func deleteUserStorage(
	userID string,
) error {
	ctx := context.Background()
	bucketName := "cheers-a275e.appspot.com"

	// Construct the path to the folder
	folderPath := fmt.Sprintf("%s/", userID)

	// Create a new Google Cloud Storage client
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	// Get a handle to the storage bucket
	bucket := client.Bucket(bucketName)

	// List all objects in the bucket with the given prefix (i.e. the folder)
	it := bucket.Objects(ctx, &storage.Query{Prefix: folderPath})
	for {
		// Get the next object
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		// Delete the object
		if err := bucket.Object(attrs.Name).Delete(ctx); err != nil {
			return err
		}
	}

	return nil
}
