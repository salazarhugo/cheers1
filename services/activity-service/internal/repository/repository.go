package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
)

type ActivityRepository interface {
	CreateActivity(userID string, activity *activity.Activity) error
	ListActivity(userID string) ([]*activity.Activity, error)
}

type activityRepository struct {
}

func NewRepository() ActivityRepository {
	return &activityRepository{}
}
