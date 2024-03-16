package redisdb

func (cache *redisCache) IsAdmin(
	userID string,
	roomID string,
) (bool, error) {
	room, err := cache.GetRoomWithId(userID, roomID)
	if err != nil {
		return false, err
	}

	var isAdmin = false
	for _, admin := range room.Admins {
		if admin == userID {
			isAdmin = true
			break
		}
	}

	return isAdmin, nil
}
