package main

import (
	"context"
	v1 "github.com/salazarhugo/cheers1/proto/out/cheers/api/v1"
)

func (s *out.MainApiServiceServer) GetUser(
	ctx context.Context,
	request *v1.GetUserRequest,
) (*v1.GetUserResponse, error) {
	return &v1.GetUserResponse{
		PostCount:      99,
		FollowBack:     false,
		FollowingCount: 0,
		FollowersCount: 0,
		StoryState:     "",
	}, nil
}
