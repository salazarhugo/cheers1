package user_service

import (
	"encoding/json"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"log"
	"testing"
)

func TestGetPostItem(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()

	response, err := repo.GetPostItem(
		"user1",
		"afa0af7b-02bf-4af1-8b18-a653ac4db492",

	)
	if err != nil {
		t.Error("failed to get feed: ", err)
		return
	}
	indent, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		return
	}
	log.Println(string(indent))
}
