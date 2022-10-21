package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
)

func (s *Server) UpdateStory(
	ctx context.Context,
	request *pb.UpdateStoryRequest,
) (*pb.StoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
