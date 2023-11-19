package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
)

func (p *userRepository) GetUser(
	userID string,
	otherUserID string,
) (*pb.GetUserResponse, error) {

	response, err := p.GetUserById(otherUserID)
	if err != nil {
		return nil, err
	}

	user := &pb.GetUserResponse{
		User:               response.ToUserPb(),
		PostCount:          0,
		FollowersCount:     0,
		FollowingCount:     0,
		HasFollowed:        false,
		StoryState:         0,
		Friend:             false,
		Requested:          false,
		HasRequestedViewer: false,
		FriendsCount:       0,
	}

	return user, nil
}
