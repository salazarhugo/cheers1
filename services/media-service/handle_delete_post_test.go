package user_service

import (
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"testing"
)

func TestHandleDeletePost(t *testing.T) {
	postID := "fac3b0a4-5be5-4fa8-9921-cfd124b7fa4a"

	err := repository.HandlePostDelete(postID)
	if err != nil {
		t.Error(err)
	}
}
