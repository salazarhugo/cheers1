package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListRoom(
	ctx context.Context,
	request *pb.ListRoomRequest,
) (*pb.ListRoomResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	rooms, err := s.chatRepository.ListRoom(userID)
	if err != nil {
		return nil, err
	}

	return &pb.ListRoomResponse{Rooms: rooms}, nil
}
