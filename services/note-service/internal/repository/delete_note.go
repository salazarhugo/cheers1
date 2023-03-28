package repository

import (
	"context"
)

func (r repository) DeleteNote(
	userID string,
) error {
	ctx := context.Background()
	_, err := r.firestore.Collection("notes").Doc(userID).Delete(ctx)

	return err
}
