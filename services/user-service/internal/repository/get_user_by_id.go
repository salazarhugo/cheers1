package repository

func (p *userRepository) GetUserById(
	userID string,
) (User, error) {
	db := p.spanner
	var user User

	result := db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
