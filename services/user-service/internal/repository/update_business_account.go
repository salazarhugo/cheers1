package repository

func (p *userRepository) UpdateBusinessAccount(
	userID string,
	isBusinessAccount bool,
) error {
	_, err := p.GetUserById(userID)
	if err != nil {
		return err
	}

	//user.User.IsBusinessAccount = isBusinessAccount
	//err = p.UpdateUser(userID, user.User)
	//if err != nil {
	//	return err
	//}
	//
	return nil
}
