package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salazarhugo/cheers1/services/postservice/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*emptypb.Empty, error) {
	_, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyID := request.GetId()
	if partyID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be empty")
	}

	err = s.partyRepository.DeletePost(partyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return &empty.Empty{}, nil
}
