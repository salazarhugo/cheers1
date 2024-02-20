package user_service

import (
	post "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"testing"
)

func TestHandleCreatePost(t *testing.T) {
	post := &post.Post{
		Id: "1f9e2762-d51e-48e2-ab3b-56eafca63b29",
		PostMedia: []*post.PostMedia{
			&post.PostMedia{
				AccessibilityCaption: "",
				MediaType:            0,
				ImageVersions: []*post.ImageVersion{
					&post.ImageVersion{
						Url:    "https://storage.googleapis.com/download/storage/v1/b/cheers-posts/o/1708447679555?generation=1708447680034132&alt=media",
						Ref:    "gs://cheers-posts/1708447679555",
						Width:  0,
						Height: 0,
					},
				},
			},
		},
	}

	err := repository.HandlePostCreate(post)
	if err != nil {
		t.Error(err)
	}
}
