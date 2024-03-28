package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendReadReceipt(
	ctx context.Context,
	request *chat.SendReadReceiptRequest,
) (*chat.Empty, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.chatRepository.SendReadReceipt(userID, request.RoomId)
	if err != nil {
		return nil, err
	}

	return &chat.Empty{}, nil
}
