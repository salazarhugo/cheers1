package mapper

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
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
	}
}
