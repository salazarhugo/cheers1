package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/note-service/internal/domain"
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

	params := domain.CreateNoteUseCaseParams{
		UserID:   userID,
		NoteType: request.GetType(),
		Text:     request.GetText(),
		DrinkID:  request.GetDrinkId(),
	}

	_, err = s.domain.CreateNoteUseCase(params)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	noteItem, err := s.repository.GetNoteItem(userID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &note.CreateNoteResponse{
		Note: noteItem.ToNotePb(),
	}, nil
}
