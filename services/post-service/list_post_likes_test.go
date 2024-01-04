package user_service

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"log"
	"testing"
)

func TestListPostLikes(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()
	postID := "9669eebb-4ac3-4df8-916a-61a1d6a3ce62"

	var err error

	// Create channels to receive results
	usersChan := make(chan []*user.UserItem, 1)
	totalChan := make(chan int64, 1)
	errChan := make(chan error, 2)

	go func() {
		users, err := repo.ListPostLikes(
			postID,
			1,
			10,
		)
		errChan <- err
		usersChan <- users
	}()

	go func() {
		total, err := repo.GetPostTotalLikes(
			postID,
		)
		errChan <- err
		totalChan <- total
	}()

	// Wait for both goroutines to finish
	users := <-usersChan
	total := <-totalChan
	err = <-errChan

	if err != nil {
		t.Error(err)
	}
	log.Println(users)
	log.Println(total)
}
