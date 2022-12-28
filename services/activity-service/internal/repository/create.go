package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (a activityRepository) CreateActivity(
	userID string,
	activity *activity.Activity,
) error {
	ctx := context.Background()

	client, err := utils.InitializeAppDefault().Firestore(ctx)
	if err != nil {
		return err
	}

	doc := client.Collection("users").Doc(userID).Collection("activity").NewDoc()
	activity.Id = doc.ID

	m, err := utils.ProtoToMap(activity)
	_, err = doc.Set(ctx, m)
	if err != nil {
		return err
	}

	return nil
}
