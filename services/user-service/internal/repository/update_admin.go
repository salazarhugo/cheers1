package repository

func (p *userRepository) UpdateAdmin(
	userID string,
	isAdmin bool,
) error {
	user, err := p.GetUserById(userID)
	if err != nil {
		return err
	}

	user.IsAdmin = isAdmin
	_, err = p.UpdateUser(&user)
	if err != nil {
		return err
	}

	return nil
}
