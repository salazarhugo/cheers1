package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListPostLikes(
	ctx context.Context, request *pb.ListPostLikesRequest,
) (*pb.ListPostLikesResponse, error) {
	_, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving viewerID")
	}

	// Create channels to receive results
	usersChan := make(chan []*user.UserItem, 1)
	totalChan := make(chan int64, 1)
	errChan := make(chan error, 2)

	// Run ListPostLikes concurrently
	go func() {
		users, err := s.postRepository.ListPostLikes(
			request.GetPostId(),
			int(request.GetPage()),
			int(request.GetPageSize()),
		)
		errChan <- err
		usersChan <- users
	}()

	// Run GetPostTotalLikes concurrently
	go func() {
		total, err := s.postRepository.GetPostTotalLikes(
			request.GetPostId(),
		)
		errChan <- err
		totalChan <- total
	}()

	// Wait for both goroutines to finish
	users := <-usersChan
	total := <-totalChan
	err = <-errChan

	if err != nil {
		return nil, err
	}

	return &pb.ListPostLikesResponse{
		Users:     users,
		LikeTotal: total,
	}, nil
}
