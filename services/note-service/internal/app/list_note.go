package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListNote(
	ctx context.Context,
	request *note.ListFriendNoteRequest,
) (*note.ListFriendNoteResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	items, err := s.repository.ListNote(
		userID,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &note.ListFriendNoteResponse{
		Items: items,
	}, nil
}
