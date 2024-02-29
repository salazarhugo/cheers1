package mapper

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"strings"
)

type UserStatusItem models.UserStatusItem

// ToNotePb Domain -> Protobuf
func (u UserStatusItem) ToNotePb() *note.Note {
	return &note.Note{
		UserId:   u.UserId,
		Text:     u.Text,
		Name:     u.Name,
		Username: u.Username,
		Picture:  u.Picture,
		Created:  u.CreatedAt.Unix(),
		Type:     ParseNoteType(u.Type),
		Drink: &drink.Drink{
			Id:   u.DrinkId.StringVal,
			Name: u.Name,
			Icon: u.DrinkIcon,
		},
	}
}

func ParseNoteType(s string) note.NoteType {
	switch strings.ToUpper(s) {
	case "DRINKING":
		return note.NoteType_DRINKING
	case "SEARCHING":
		return note.NoteType_SEARCHING
	default:
		return note.NoteType_NOTHING
	}
}
