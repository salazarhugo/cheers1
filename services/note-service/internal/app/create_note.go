package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNote(
	ctx context.Context,
	request *note.CreateNoteRequest,
) (*note.CreateNoteResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	if request.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty field: text")
	}

	if len(request.GetText()) > 60 {
		return nil, status.Error(codes.InvalidArgument, "text must be less or equal to 60 characters")
	}

	noteID, err := s.repository.CreateNote(
		userID,
		request.GetText(),
	)

	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	res, err := s.repository.GetNote(noteID)

	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &note.CreateNoteResponse{
		Note: res,
	}, nil
}
