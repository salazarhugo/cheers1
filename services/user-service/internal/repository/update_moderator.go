package repository

func (p *userRepository) UpdateModerator(
	userID string,
	isModerator bool,
) error {
	_, err := p.GetUser(userID, userID)
	if err != nil {
		return err
	}

	//user.User.IsModerator = isModerator
	//err = p.UpdateUser(userID, user.User)
	//if err != nil {
	//	return err
	//}

	return nil
}
