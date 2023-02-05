package repository

func (p *userRepository) UpdateAdmin(
	userID string,
	isAdmin bool,
) error {
	user, err := p.GetUser(userID, userID)
	if err != nil {
		return err
	}

	user.User.IsAdmin = isAdmin
	err = p.UpdateUser(userID, user.User)
	if err != nil {
		return err
	}

	return nil
}
