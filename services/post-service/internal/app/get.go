package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPost(
	ctx context.Context,
	request *pb.GetPostRequest,
) (*pb.PostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	postID := request.GetId()

	post, err := s.partyRepository.GetPost(userID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get post")
	}

	return post, nil
}
