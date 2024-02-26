package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
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

	GetUserById(
		userID string,
	) (models.User, error)

	GetUserByUsername(
		username string,
	) (*models.User, error)

	GetUserWithViewer(
		viewerID string,
		otherUserID string,
	) (*pb.GetUserResponse, error)

	UpdateUser(
		user *models.User,
	) (string, error)

	DeleteUser(
		userID string,
	) error

	UpdateBusinessAccount(
		userID string,
		isBusinessAccount bool,
	) error

	UpdateAdmin(
		userID string,
		isAdmin bool,
	) error

	UpdateModerator(
		userID string,
		isModerator bool,
	) error

	VerifyUser(
		userID string,
		verified bool,
	) error

	GetUsersIn(
		userIDs []string,
	) ([]*models.User, error)

	GetUserItemsIn(
		userID string,
		userIDs []string,
	) ([]*user.UserItem, error)

	SearchUser(
		query string,
	) ([]*models.User, error)

	ListSuggestions(
		userID string,
	) ([]*user.UserItem, error)

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
