package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteStory(
	ctx context.Context,
	request *pb.DeleteStoryRequest,
) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
