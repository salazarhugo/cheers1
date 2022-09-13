package main

import (
	"context"
	"fmt"
	"salazar/cheers/user/proto/userpb"
	"sync"
)

func newServer() *userServiceServer {
	s := &userServiceServer{}
	fmt.Println(s)
	return s
}

type userServiceServer struct {
	userpb.UnimplementedUserServiceServer
	mu sync.Mutex
}

func (s *userServiceServer) GetUser(
	ctx context.Context,
	request *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {

	user := &userpb.GetUserResponse{
		User:           nil,
		PostCount:      0,
		FollowBack:     false,
		FollowingCount: 0,
		FollowersCount: 0,
		StoryState:     "",
	}

	return user, nil
}
