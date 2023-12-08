package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/post-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) FeedPost(
	ctx context.Context,
	request *pb.FeedPostRequest,
) (*pb.FeedPostResponse, error) {
	userId, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	friendIds, err := services.GetFriendList(userId)
	if err != nil {
		return nil, err
	}

	log.Println(friendIds)

	posts, err := s.postRepository.FeedPost(
		append(friendIds, userId),
		int(request.Page),
		int(request.PageSize),
	)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
