package repository

func (a *authRepository) GetUserByUsername(
	username string,
) (*User, error) {
	db := a.spanner
	var user User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
