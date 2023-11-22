package repository

func (a authRepository) CreateCredential(
	credential *Credential,
) error {
	db := a.spanner

	result := db.Create(&credential)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
