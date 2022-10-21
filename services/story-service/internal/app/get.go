package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
)

func (s *Server) GetStory(
	ctx context.Context,
	request *pb.GetStoryRequest,
) (*pb.StoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
