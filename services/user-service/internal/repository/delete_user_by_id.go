package repository

func (p *userRepository) DeleteUserById(
	userID string,
) error {
	db := p.spanner
	result := db.Delete(&User{ID: userID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
