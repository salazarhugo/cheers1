package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/location/v1"
	"github.com/salazarhugo/cheers1/services/location-service/internal/service"
	"log"
	"strconv"
)

func (r repository) ListFriendLocation(
	ctx context.Context,
	userId string,
) ([]*location.Location, error) {
	friends, err := service.ListFriends(ctx, userId)
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

	items := make([]*location.Location, 0)
	for i, loc := range locations {
		if loc == nil {
			continue
		}
		friend := friends[i]

		ghostMode, err := r.GetGhostMode(friend.Id)
		if err != nil {
			log.Println(err)
		}

		if ghostMode == true {
			continue
		}

		// Get the last updated timestamp
		str, err := r.redis.HGet(
			ctx,
			keyLastUpdated,
			friend.Id,
		).Result()
		if err != nil {
			log.Println(err)
		}

		// Convert the timestamp string to an int64
		lastUpdated, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Println(err)
			lastUpdated = 0
		}

		items = append(items, &location.Location{
			Latitude:    loc.Latitude,
			Longitude:   loc.Longitude,
			UserId:      friend.Id,
			Name:        friend.Name,
			Username:    friend.Username,
			Picture:     friend.Picture,
			Verified:    friend.Verified,
			LastUpdated: lastUpdated,
		})
	}

	return items, nil
}
