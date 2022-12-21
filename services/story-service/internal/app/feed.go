package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) FeedStory(
	ctx context.Context,
	request *pb.FeedStoryRequest,
) (*pb.FeedStoryResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	log.Print(request.String())
	posts, err := s.storyRepository.FeedStory(userID, request)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
