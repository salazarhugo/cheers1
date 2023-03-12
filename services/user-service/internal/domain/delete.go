package domain

func (userService *UserService) DeleteUser(
	userID string,
) error {
	_, err := userService.userRepository.GetUserNode(userID)
	if err != nil {
		return err
	}

	err = userService.userRepository.DeleteUser(userID)
	if err != nil {
		return err
	}

	err = deleteUserStorage(userID)
	if err != nil {
		return err
	}

	err = deleteUserDocument(userID)
	if err != nil {
		return err
	}

	return nil
}
