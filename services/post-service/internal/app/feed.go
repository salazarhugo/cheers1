package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) FeedPost(ctx context.Context, request *pb.FeedPostRequest) (*pb.FeedPostResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	log.Print(request)
	posts, err := s.postRepository.FeedPost(userID, request)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
