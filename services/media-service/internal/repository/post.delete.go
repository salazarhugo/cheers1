package repository

import (
	post "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
)

func HandlePostDelete(post *post.Post) error {
	repo := NewPostRepository()

	postMedias, err := repo.ListPostMedia(post.GetId())
	if err != nil {
		return err
	}

	for _, media := range postMedias {
		// Parse image reference
		bucketName, path, err := ParseStorageRef(media.Ref)
		if err != nil {
			return err
		}

		// Delete image from cloud storage
		err = DeleteCloudStorageObject(bucketName, path)
		if err != nil {
			return err
		}

		// Delete image metadata
		err = repo.DeletePostMedia(media.PostMediaId)
		if err != nil {
			return err
		}
	}

	return nil
}
