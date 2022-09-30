package main

import v1 "github.com/salazarhugo/cheers1/proto/out/cheers/api/v1"

func (s *mainApiServiceServer) GetUser(
	ctx context.Context,
	request *v1.GetUserRequest,
) (*v1.GetUserResponse, error) {
	return user, nil
}
