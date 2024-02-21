package repository

import (
	"context"
	"fmt"
)

func HandlePostCreate(postID string) error {
	ctx := context.Background()
	repo := NewPostRepository()

	postMedias, err := repo.ListPostMedia(postID)
	if err != nil {
		return err
	}

	for _, media := range postMedias {
		// Parse image reference
		bucketName, path, err := ParseStorageRef(media.Ref)
		if err != nil {
			return err
		}

		// Get image from cloud storage
		img, err := GetCloudStorageImage(bucketName, path)
		if err != nil {
			return err
		}

		// Resize the image to a 1080x1080 square
		resizedImg, err := SquareResizeImage(*img)
		if err != nil {
			return fmt.Errorf("SquareResizeImage: %v", err)
		}

		// Upload the resized image back to Cloud Storage
		object, err := UploadImageToCloudStorage(ctx, bucketName, path, resizedImg)
		if err != nil {
			return fmt.Errorf("UploadImageToCloudStorage: %v", err)
		}

		objectAttributes, err := object.Attrs(ctx)
		if err != nil {
			return fmt.Errorf("failed getting object attributes: %v", err)
		}

		// Update download link
		media.Url = objectAttributes.MediaLink
		err = repo.UpdatePostMedia(media)
		if err != nil {
			return err
		}
	}

	return nil
}
