package repository

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	service "github.com/salazarhugo/cheers1/services/media-service/internal/services"
	"image"
	"image/jpeg"
)

func HandlePostCreate(postID string) error {
	ctx := context.Background()
	post, err := service.GetPost(postID)
	if err != nil {
		return err
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	for _, media := range post.PostMedia {
		originalVersion := media.ImageVersions[0]
		bucketName, path, err := ParseStorageRef(originalVersion.GetRef())
		if err != nil {
			return err
		}

		rc, err := client.Bucket(bucketName).Object(path).NewReader(ctx)
		if err != nil {
			return err
		}
		defer rc.Close()

		// Decode the image
		img, _, err := image.Decode(rc)
		if err != nil {
			return fmt.Errorf("image.Decode: %v", err)
		}

		// Resize the image to 1080x1080
		resizedImg, err := ResizeImage(img, 1080, 1080)
		if err != nil {
			return fmt.Errorf("ResizeImage: %v", err)
		}

		// Encode the resized image as JPEG
		var buf bytes.Buffer
		err = jpeg.Encode(&buf, resizedImg, nil)
		if err != nil {
			return fmt.Errorf("jpeg.Encode: %v", err)
		}

		// Upload the resized image back to Cloud Storage
		//resizedObjectName := "resized_" + objectName
		//err = UploadToCloudStorage(ctx, bucketName, resizedObjectName, buf.Bytes())
		//if err != nil {
		//	return fmt.Errorf("UploadToCloudStorage: %v", err)
		//}

	}

	return nil
}
