package app

import (
	"context"
	story "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ViewStory(
	ctx context.Context,
	request *story.ViewStoryRequest,
) (*story.ViewStoryResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	storyID := request.GetId()
	if storyID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be blank")
	}

	response, err := s.storyRepository.ViewStory(userID, storyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create story")
	}

	return response, nil
}
