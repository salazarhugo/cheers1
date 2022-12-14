package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetInbox(
	ctx context.Context,
	request *chat.GetInboxRequest,
) (*chat.GetInboxResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	inbox, err := s.chatRepository.GetInbox(userID)
	if err != nil {
		return nil, err
	}

	return &chat.GetInboxResponse{Inbox: inbox}, nil
}
