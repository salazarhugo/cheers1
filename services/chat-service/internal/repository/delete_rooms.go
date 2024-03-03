package repository

func (c chatRepository) DeleteRooms(
	userID string,
) error {
	rooms, err := c.cache.ListRooms(userID)
	if err != nil {
		return err
	}

	for _, room := range rooms {
		err := c.DeleteRoom(userID, room.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
