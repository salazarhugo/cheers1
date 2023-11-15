package repository

func (p *userRepository) GetUserByUsername(
	username string,
) (User, error) {
	db := p.spanner
	var user User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
