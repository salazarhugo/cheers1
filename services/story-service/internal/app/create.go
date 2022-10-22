package app

import (
	"context"
	story "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateStory(
	ctx context.Context,
	request *story.CreateStoryRequest,
) (*story.StoryResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	storyReq := request.GetStory()
	if storyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "story parameter can't be nil")
	}

	storyID, err := s.storyRepository.CreateStory(userID, storyReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create story")
	}

	story, err := s.storyRepository.GetStory(userID, storyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get story")
	}

	return story, nil
}
