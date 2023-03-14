package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

const (
	keyUser          = "user"
	keyUserLocations = "locations"
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

	return err
}
