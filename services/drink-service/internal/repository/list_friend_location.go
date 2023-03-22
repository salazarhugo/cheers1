package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/services/drink-service/internal/service"
	"log"
	"strconv"
)

func (r repository) ListDrink(
	userId string,
) ([]*drink.Drink, error) {
	ctx := context.Background()

	friends, err := service.List(userId)
	if err != nil {
		return nil, err
	}

	// list of friend IDs
	friendIDs := make([]string, len(friends))
	for i, friend := range friends {
		friendIDs[i] = friend.Id
	}

	// Use the GeoPos command to get the drinks of all the users in the friendIDs list
	drinks, err := r.redis.GeoPos(
		ctx,
		keyUserDrinks,
		friendIDs...,
	).Result()

	if err != nil {
		return nil, err
	}

	items := make([]*drink.Drink, 0)
	for i, loc := range drinks {
		if loc == nil {
			continue
		}
		friend := friends[i]

		// Get the last updated timestamp
		str, err := r.redis.HGet(ctx, keyLastUpdated,
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

		items = append(items, &drink.Drink{
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
