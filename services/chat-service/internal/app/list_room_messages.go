package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"log"
)

func (s *Server) ListRoomMessages(
	ctx context.Context,
	req *chat.ListRoomMessagesRequest,
) (*chat.ListRoomMessagesResponse, error) {
	messages, err := s.chatRepository.ListRoomMessages(req.RoomId, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &chat.ListRoomMessagesResponse{Messages: messages}, nil
}
