package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UnlikePost(
	ctx context.Context,
	request *pb.UnlikePostRequest,
) (*pb.UnlikePostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	response, err := s.partyRepository.UnlikePost(userID, request.Id)
	if err != nil {
		return nil, err
	}

	return response, nil
}
