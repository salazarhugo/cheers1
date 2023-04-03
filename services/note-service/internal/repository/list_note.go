package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"github.com/salazarhugo/cheers1/services/note-service/internal/service"
	"google.golang.org/api/iterator"
	"log"
)

func (r repository) ListNote(
	userID string,
) ([]*note.Note, error) {
	ctx := context.Background()

	friends, err := service.ListFriends(userID)
	if err != nil {
		return nil, err
	}

	// list of friend IDs
	friendIDs := make([]string, len(friends))
	for i, friend := range friends {
		friendIDs[i] = friend.Id
	}
	// Add current user
	friendIDs = append(friendIDs, userID)

	documents := r.firestore.Collection("notes").Where("userId", "in", friendIDs).Documents(ctx)

	notes := make([]*note.Note, 0)

	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
		}
		item := &note.Note{}
		m := doc.Data()
		err = mapper.MapToProto(item, m)
		if err != nil {
			log.Println(err)
		}
		notes = append(notes, item)
	}

	return notes, nil
}
