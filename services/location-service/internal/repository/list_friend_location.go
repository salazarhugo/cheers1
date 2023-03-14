package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/location/v1"
	"github.com/salazarhugo/cheers1/services/location-service/internal/service"
)

func (r repository) ListFriendLocation(
	userId string,
) ([]*location.Location, error) {
	ctx := context.Background()

	friends, err := service.ListFriends(userId)
	if err != nil {
		return nil, err
	}

	// list of friend IDs
	friendIDs := make([]string, len(friends))
	for i, friend := range friends {
		friendIDs[i] = friend.Id
	}

	// Use the GeoPos command to get the locations of all the users in the friendIDs list
	locations, err := r.redis.GeoPos(
		ctx,
		keyUserLocations,
		friendIDs...,
	).Result()

	if err != nil {
		return nil, err
	}

	items := make([]*location.Location, len(friends))
	for i, loc := range locations {
		friend := friends[i]
		items[i] = &location.Location{
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			UserId:    friend.Id,
			Name:      friend.Name,
			Username:  friend.Username,
			Picture:   friend.Picture,
			Verified:  friend.Verified,
		}
	}

	return items, nil
}
