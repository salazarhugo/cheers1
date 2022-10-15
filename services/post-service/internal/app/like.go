package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LikePost(
	ctx context.Context,
	request *pb.LikePostRequest,
) (*pb.LikePostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	response, err := s.partyRepository.LikePost(userID, request.Id)
	if err != nil {
		return nil, err
	}

	return response, nil
}
