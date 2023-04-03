package repository

import (
	"context"
	note2 "github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"github.com/salazarhugo/cheers1/services/note-service/internal/service"
	"time"
)

func (r repository) CreateNote(
	creatorID string,
	text string,
) (string, error) {
	ctx := context.Background()
	doc := r.firestore.Collection("notes").Doc(creatorID)

	user, err := service.GetUser(creatorID)
	if err != nil {
		return "", err
	}

	note := &note2.Note{
		UserId:   creatorID,
		Text:     text,
		Name:     user.Name,
		Username: user.Username,
		Picture:  user.Picture,
		Created:  time.Now().Unix(),
	}

	m, err := mapper.ProtoToMap(note)
	m["deleteAt"] = time.Now().AddDate(0, 0, 1)
	_, err = doc.Set(ctx, m)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
