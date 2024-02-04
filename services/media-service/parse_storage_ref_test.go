package user_service

import (
	"github.com/salazarhugo/cheers1/services/image-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStorageRef(t *testing.T) {
	ref := "gs://cheers-a275e.appspot.com/Ii7mFI34JaYKYsOxYg90fYToa182/posts/01da517e-e15a-4783-bcb1-00f7b232bd5e"

	bucketName, path, err := repository.ParseStorageRef(ref)
	if err != nil {
		return
	}

	assert.Equal(t, "cheers-a275e.appspot.com", bucketName)
	assert.Equal(t, "Ii7mFI34JaYKYsOxYg90fYToa182/posts/01da517e-e15a-4783-bcb1-00f7b232bd5e", path)
}
