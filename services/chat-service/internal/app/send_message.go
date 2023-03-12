package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) SendMessage(
	ctx context.Context,
	request *chat.SendMessageRequest,
) (*chat.SendMessageResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = CheckPermissions()
	if err != nil {
		return nil, err
	}

	msg, err := s.chatRepository.SendMessage(
		request.ClientId,
		userID,
		request.RoomId,
		request.Text,
		request.GetReplyTo(),
	)

	msg.Status = chat.Message_SENT

	return &chat.SendMessageResponse{Message: msg}, nil
}

func CheckPermissions() error {
	//isMember := roomCache.IsMember(userId, msg.Room.Id)
	isMember := true
	if isMember == false {
		log.Println("Can't access rooms you are not member of")
		return status.Error(codes.PermissionDenied, "you are not member of this group chat")
	}
	return nil
}
