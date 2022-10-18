package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateChat(ctx context.Context, request *pb.CreateChatReq) (*pb.Room, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	UUIDs := append([]string{userID}, request.UserIds...)
	UUIDs = unique(UUIDs)

	if len(UUIDs) < 2 {
		return nil, status.Error(codes.InvalidArgument, "chats must have at least 2 users")
	}

	// Direct Chat
	if len(UUIDs) == 2 {
		room := roomCache.GetOrCreateDirectRoom(UUIDs[0], UUIDs[1])
		return room, nil
	}

	room := roomCache.CreateGroup(request.GroupName, UUIDs)
	return room, nil
}
