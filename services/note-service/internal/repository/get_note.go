package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (r repository) GetNote(
	userID string,
) (*note.Note, error) {
	ctx := context.Background()

	doc, err := r.firestore.Collection("notes").Where("userId", "==", userID).Documents(ctx).Next()

	item := &note.Note{}
	m := doc.Data()

	err = mapper.MapToProto(item, m)
	if err != nil {
		return nil, err
	}

	return item, nil
}
