package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(
		userID string,
		username string,
		name string,
		picture string,
		email string,
	) (string, error)

	GetUserById(userID string) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUser(userID string, otherUserID string) (*pb.GetUserResponse, error)
	UpdateUser(userID string, user *user.User) error
	DeleteUserById(userID string) error

	UpdateBusinessAccount(userID string, isBusinessAccount bool) error
	UpdateAdmin(userID string, isAdmin bool) error
	UpdateModerator(userID string, isModerator bool) error
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

	ListSuggestions(
		userID string,
	) ([]*user.UserItem, error)

	CreateFriend(
		from string,
		to string,
	) error

	CreateFriendRequest(
		from string,
		to string,
	) error

	DeleteFriendRequest(
		from string,
		to string,
	) error

	DeleteFriend(
		from string,
		to string,
	) error

	CheckUsername(
		username string,
	) (bool, error)
}

type userRepository struct {
	spanner *gorm.DB
	driver  neo4j.Driver
}

func NewUserRepository() UserRepository {
	return &userRepository{
		spanner: utils.GetSpanner(),
		driver:  utils.GetDriver(),
	}
}
