package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"
)

const (
	keyUser          = "user"
	keyUserLocations = "locations"
	keyLastUpdated   = "updated"
)

func getKeyUserLocations(userId string) string {
	return fmt.Sprintf("%s:%s", userId, userId)
}

func (r repository) UpdateLocation(
	userId string,
	latitude float64,
	longitude float64,
) error {
	ctx := context.Background()

	// Use the GeoAdd command to add the user's location to the "user_locations" key
	err := r.redis.GeoAdd(ctx, keyUserLocations, &redis.GeoLocation{
		Name:      userId,
		Latitude:  latitude,
		Longitude: longitude,
	}).Err()

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// Set last updated timestamp
	err = r.redis.HSet(ctx, keyLastUpdated,
		userId,
		timestamp,
	).Err()

	return err
}
