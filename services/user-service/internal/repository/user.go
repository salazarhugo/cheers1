package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
)

type UserRepository interface {
	CreateUser(userID string, post *user.User) (string, error)
	GetUser(userID string, postID string) (*pb.GetUserResponse, error)
	UpdateUser(user *user.User) (*user.User, error)
	DeleteUser(id string) error

	FollowUser(userID string, otherUserID string) error
	UnfollowUser(userID string, otherUserID string) error

	BlockUser(userID string, otherUserID string) error
	UnblockUser(userID string, otherUserID string) error

	SearchUser(userID string, query string) ([]*user.User, error)
}

type postRepository struct {
	driver neo4j.Driver
}

func NewUserRepository(driver neo4j.Driver) UserRepository {
	return &postRepository{driver: driver}
}
