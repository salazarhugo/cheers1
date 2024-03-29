package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFriendNote(
	ctx context.Context,
	request *note.ListFriendNoteRequest,
) (*note.ListFriendNoteResponse, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve viewerID")
	}

	items, err := s.repository.FeedNote(
		viewerID,
		1,
		10,
	)

	if err != nil {
		return nil, err
	}

	return &note.ListFriendNoteResponse{
		Items: items,
	}, nil
}
