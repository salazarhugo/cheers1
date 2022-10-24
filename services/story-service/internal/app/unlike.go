package app

import (
	"context"
	story "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UnlikeStory(
	ctx context.Context,
	request *story.UnlikeStoryRequest,
) (*story.UnlikeStoryResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	storyID := request.GetId()
	if storyID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be blank")
	}

	response, err := s.storyRepository.UnlikeStory(userID, storyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to unlike story")
	}

	return response, nil
}
