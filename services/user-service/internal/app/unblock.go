package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UnblockUser(
	ctx context.Context,
	request *pb.UnblockUserRequest,
) (*empty.Empty, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	err = s.userRepository.UnblockUser(userID, request.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
