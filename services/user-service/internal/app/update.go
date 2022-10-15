package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
)

func (s *Server) UpdatePost(
	ctx context.Context,
	request *pb.UpdatePostRequest,
) (*pb.PostResponse, error) {
	panic("implement me")
}
