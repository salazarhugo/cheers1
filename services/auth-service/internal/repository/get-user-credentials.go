package repository

func (a *authRepository) GetUserCredentials(
	username string,
) ([]*Credential, error) {
	db := a.spanner
	var credentials []*Credential

	result := db.Table("users").Select("credentials.*").Joins("JOIN credentials ON users.authn_id = credentials.user_id", username).Where("username = ?", username).Scan(&credentials)
	if result.Error != nil {
		return nil, result.Error
	}

	return credentials, nil
}
