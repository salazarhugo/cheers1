package user_service

import (
	"github.com/salazarhugo/cheers1/services/media-service/internal/models"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"log"
	"testing"
)

func TestCreateMedia(t *testing.T) {
	repo := repository.NewPostRepository()
	postMediaID, err := repo.CreateMedia(&models.Media{
		Url:                  "https://www.google.com",
		Ref:                  "gs://cheers.com/posts/sfkewfwfw",
		Type:                 "image/jpeg",
		AccessibilityCaption: "",
	})
	if err != nil {
		t.Error(err)
	}

	log.Println(postMediaID)
}
