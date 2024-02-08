package user_service

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"log"
	"testing"
)

func TestGetPostItem(t *testing.T) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
	}

	ref := "gs://cheers-a275e.appspot.com/Ii7mFI34JaYKYsOxYg90fYToa182/posts/01da517e-e15a-4783-bcb1-00f7b232bd5e"
	bucketName, path, err := repository.ParseStorageRef(ref)
	if err != nil {
		return
	}

	defer client.Close()
	_, err = client.Bucket(bucketName).Object(path).Attrs(ctx)
	if err != nil {
		log.Println(err)
	}
}
