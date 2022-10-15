package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*emptypb.Empty, error) {
	_, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	id := request.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be empty")
	}

	err = s.userRepository.DeleteUser(id)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed delete user")
	}

	return &empty.Empty{}, nil
}
