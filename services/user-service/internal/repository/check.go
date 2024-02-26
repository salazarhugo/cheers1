package repository

func (p *userRepository) CheckUsername(
	username string,
) (bool, error) {
	db := p.spanner

	var exists bool

	err := db.
		Table("users").
		Select("EXISTS (SELECT 1 FROM users WHERE username = ?)", username).
		Scan(&exists).
		Error
	if err != nil {
		return false, err
	}

	return exists, err
}
