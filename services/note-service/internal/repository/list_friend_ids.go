package repository

func (r *repository) ListFriendIds(
	userID string,
) ([]string, error) {
	db := r.spanner

	friendIDs := make([]string, 0)

	err := db.
		Table("friendships").
		Select("user_id1").
		Where("user_id2 = ?", userID).
		Scan(&friendIDs).
		Error
	if err != nil {
		return nil, err
	}

	return friendIDs, nil
}
