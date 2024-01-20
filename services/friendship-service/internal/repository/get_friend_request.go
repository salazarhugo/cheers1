package repository

func (r repository) GetFriendRequest(
	userID string,
	user2ID string,
) (*FriendRequest, error) {
	db := r.spanner

	var friendRequest FriendRequest

	err := db.Table("friend_requests").
		Where("user_id1 = ? AND user_id2 = ?", userID, user2ID).
		First(&friendRequest).
		Error

	if err != nil {
		return nil, err
	}

	return &friendRequest, nil
}
