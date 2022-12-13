package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) DeleteRoom(
	ctx context.Context,
	request *chat.DeleteRoomRequest,
) (*chat.DeleteRoomResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.chatRepository.DeleteRoom(userID, request.RoomId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &chat.DeleteRoomResponse{}, nil
}
