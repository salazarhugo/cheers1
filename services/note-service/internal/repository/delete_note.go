package repository

func (r *repository) DeleteNote(
	userID string,
) error {
	db := r.spanner

	err := db.
		Exec("DELETE FROM user_status WHERE UserId = ?", userID).
		Error

	return err
}
