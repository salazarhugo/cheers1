package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRoom(
	ctx context.Context,
	request *pb.CreateRoomRequest,
) (*pb.CreateRoomResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	if len(request.RecipientUsers) < 1 {
		return nil, status.Error(codes.InvalidArgument, "empty parameter recipient_users")
	}

	if request.RecipientUsers[0] == userID {
		return nil, status.Error(codes.InvalidArgument, "cannot create room with yourself")
	}

	members := append([]string{userID}, request.RecipientUsers...)

	room, err := s.chatRepository.CreateRoom(request.GroupName, members)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{
		Room: room,
	}, nil
}
