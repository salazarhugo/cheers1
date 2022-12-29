package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

type UserRepository interface {
	CreateUser(userID string, user *user.User) (string, error)
	GetUser(userID string, otherUserID string) (*pb.GetUserResponse, error)
	UpdateUser(userID string, user *pb.UpdateUserRequest) error
	DeleteUser(id string) error

	VerifyUser(userID string) error
	UnVerifyUser(userID string) error
	GetUsersIn(userIDs []string) ([]*user.User, error)
	GetUserItemsIn(userID string, userIDs []string) ([]*user.UserItem, error)
	FollowUser(userID string, otherUser string) error
	UnfollowUser(userID string, otherUserID string) error

	BlockUser(userID string, otherUserID string) error
	UnblockUser(userID string, otherUserID string) error

	SearchUser(userID string, query string) ([]*user.UserItem, error)
	ListFollowers(userID string, request *pb.ListFollowersRequest) ([]*user.UserItem, error)
	ListFollowing(userID string, request *pb.ListFollowingRequest) ([]*user.UserItem, error)
}

type userRepository struct {
	driver neo4j.Driver
}

func NewUserRepository() UserRepository {
	return &userRepository{driver: utils.GetDriver()}
}
