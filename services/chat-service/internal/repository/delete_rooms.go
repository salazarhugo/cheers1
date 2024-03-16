package repository

func (c chatRepository) DeleteRooms(
	userID string,
) error {
	chatIDs, err := c.cache.ListChatIds(userID)
	if err != nil {
		return err
	}

	for _, chatID := range chatIDs {
		err := c.DeleteRoom(userID, chatID)
		if err != nil {
			return err
		}
	}

	return nil
}
