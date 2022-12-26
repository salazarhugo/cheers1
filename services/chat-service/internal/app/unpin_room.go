package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (s *Server) UnPinRoom(
	ctx context.Context,
	req *chat.UnPinRoomRequest,
) (*chat.UnPinRoomResponse, error) {
	return &chat.UnPinRoomResponse{}, nil
}
