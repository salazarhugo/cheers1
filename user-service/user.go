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

	switch identification := request.Identification.(type) {
	case *userpb.GetUserRequest_UserId:
		fmt.Printf("Get user with userId: %s\n", identification.UserId)
	case *userpb.GetUserRequest_Username:
		fmt.Printf("Get user with username: %s\n", identification.Username)
	default:
		fmt.Println("No matching operations")
	}

	//user2.GetUser()
	user := &userpb.GetUserResponse{
		User: &userpb.User{
			Id:                 "demo",
			Name:               "Hugo Salazar",
			Email:              "hugobrock74@gmail.com",
			Verified:           false,
			Username:           "demo",
			Picture:            "",
			Bio:                "Haha",
			Website:            "",
			PhoneNumber:        "",
			Created:            0,
			RegistrationTokens: nil,
		},
		PostCount:      69,
		FollowBack:     false,
		FollowingCount: 69,
		FollowersCount: 69,
		StoryState:     "",
	}

	return user, nil
}
