package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteNote(
	ctx context.Context,
	request *note.DeleteNoteRequest,
) (*note.DeleteNoteResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.repository.DeleteNote(
		userID,
	)

	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &note.DeleteNoteResponse{}, nil
}
