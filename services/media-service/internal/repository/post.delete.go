package repository

import post "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"

func HandlePostDelete(post *post.Post) error {
	postMedias := post.GetPostMedia()

	for _, media := range postMedias {
		for _, imageVersion := range media.GetImageVersions() {
			// Parse image reference
			bucketName, path, err := ParseStorageRef(imageVersion.Ref)
			if err != nil {
				return err
			}

			// Delete image from cloud storage
			err = DeleteCloudStorageObject(bucketName, path)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
