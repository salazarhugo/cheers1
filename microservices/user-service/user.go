package main

import (
	"context"
	"fmt"
	"salazar/cheers/user/proto/out/cheers"
	cheers2 "salazar/cheers/user/proto/out/cheers/api/v1"
	"sync"
)

func newServer() *userServiceServer {
	s := &userServiceServer{}
	fmt.Println(s)
	return s
}

type userServiceServer struct {
	cheers2.UnimplementedMainServer
	mu sync.Mutex
}

func (s *userServiceServer) GetUser(
	ctx context.Context,
	request *cheers2.GetUserRequest,
) (*cheers2.GetUserResponse, error) {

	switch identification := request.Identification.(type) {
	case *cheers2.GetUserRequest_UserId:
		fmt.Printf("Get user with userId: %s\n", identification.UserId)
	case *cheers2.GetUserRequest_Username:
		fmt.Printf("Get user with username: %s\n", identification.Username)
	default:
		fmt.Println("No matching operations")
	}

	user := &cheers2.GetUserResponse{
		User: &cheers.User{
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
