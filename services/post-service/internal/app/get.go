package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
)

func (s *Server) GetPost(
	ctx context.Context,
	request *pb.GetPostRequest,
) (*pb.GetPostResponse, error) {
	postID := request.GetPostId()

	post, err := s.postRepository.GetPostById(postID)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{
		Post: post.ToPostPb(),
	}, nil
}
