package app

import (
	"context"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/services/postservice/genproto/cheers/post/v1"
)

func (s *Server) UpdatePost(ctx context.Context, request *pb.UpdatePostRequest) (*postpb.Post, error) {
	//TODO implement me
	panic("implement me")
}
