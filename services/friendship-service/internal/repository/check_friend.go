package repository

func (r repository) CheckFriend(
	user1 string,
	user2 string,
) (bool, error) {
	db := r.spanner

	var exists bool

	err := db.
		Table("friendships").
		Select("EXISTS (SELECT 1 FROM friendships WHERE user_id1 = ? AND user_id2 = ?)", user1, user2).
		Scan(&exists).
		Error
	if err != nil {
		return false, err
	}

	return exists, err
}
