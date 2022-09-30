package main

func (s *v1.mainApiServiceServer) GetUser(
	ctx context.Context,
	request *v1.GetUserRequest,
) (*v1.GetUserResponse, error) {
	return user, nil
}
