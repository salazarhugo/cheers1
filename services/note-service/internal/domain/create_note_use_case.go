package domain

import (
	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateNoteUseCaseParams struct {
	UserID   string
	NoteType note.NoteType
	Text     string
	DrinkID  *string
}

func (r *domain) CreateNoteUseCase(
	params CreateNoteUseCaseParams,
) (string, error) {

	switch params.NoteType {
	case note.NoteType_DRINKING:
		if params.DrinkID == nil {
			return "", status.Error(codes.InvalidArgument, "missing drink")
		}
	case note.NoteType_SEARCHING:
	case note.NoteType_NOTHING:
		if params.Text == "" {
			return "", status.Error(codes.InvalidArgument, "empty field: text")
		}

		if len(params.Text) > 60 {
			return "", status.Error(codes.InvalidArgument, "text must be less or equal to 60 characters")
		}

	default:
		return "", status.Error(codes.InvalidArgument, "unknown note type")
	}

	note := &models.UserStatus{
		UserId:       params.UserID,
		UserStatusId: uuid.NewString(),
		Text:         params.Text,
		Type:         params.NoteType.String(),
		DrinkId:      spanner.NullString{Valid: false},
	}

	if params.DrinkID != nil {
		note.DrinkId = spanner.NullString{StringVal: *params.DrinkID, Valid: true}
	}

	return r.repository.CreateNote(note)
}
