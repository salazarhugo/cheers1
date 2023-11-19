package domain

import (
	"context"
	"fmt"
	"github.com/salazarhugo/cheers1/libs/utils"
)

// deleteUserDocument deletes the Firestore document associated with the given user ID.
func deleteUserDocument(userID string) error {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	// Construct the path to the document
	docPath := fmt.Sprintf("users/%s", userID)

	// Delete the document
	_, err = client.Doc(docPath).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
