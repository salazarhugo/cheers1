package repository

func (a *authRepository) CreateUser(
	user *User,
) error {
	db := a.spanner
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
