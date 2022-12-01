package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListRoomMessages(
	ctx context.Context,
	req *chat.ListRoomMessagesRequest,
) (*chat.ListRoomMessagesResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	messages, err := s.chatRepository.ListRoomMessages(req.RoomId, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &chat.ListRoomMessagesResponse{Messages: messages}, nil
}
