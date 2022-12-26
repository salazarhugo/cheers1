package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BlockUser(
	ctx context.Context,
	request *pb.BlockUserRequest,
) (*pb.BlockUserResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	err = s.userRepository.BlockUser(userID, request.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.BlockUserResponse{}, nil
}
