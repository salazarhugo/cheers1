package app

import (
	"context"
	story "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
)

func (s *Server) CreateStory(
	ctx context.Context,
	request *story.CreateStoryRequest,
) (*story.StoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
