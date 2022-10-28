package repository

import (
	"context"
	"encoding/json"
	"github.com/salazarhugo/cheers1/genproto/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func (a activityRepository) ListActivity(
	userID string,
) ([]*activity.Activity, error) {
	ctx := context.Background()

	client, err := utils.InitializeAppDefault().Firestore(ctx)
	if err != nil {
		return nil, err
	}

	documents := client.Collection("users").Doc(userID).Collection("activity").Documents(ctx)

	activities := make([]*activity.Activity, 0)

	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		m := doc.Data()
		bytes, err := json.Marshal(m)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		a := &activity.Activity{}
		err = protojson.Unmarshal(bytes, a)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		activities = append(activities, a)
	}

	return activities, nil
}
