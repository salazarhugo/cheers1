package repository

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/redisdb"
	"gorm.io/gorm"
	"time"
)

type ChatRepository interface {
	CreateRoom(
		name string,
		members []string,
	) (*pb.Room, error)

	JoinRoom(
		request *pb.JoinRoomRequest,
		server pb.ChatService_JoinRoomServer,
	) error

	GetInbox(userID string) ([]*pb.RoomWithMessages, error)

	DeleteRoom(userID string, roomID string) error

	DeleteRooms(userID string) error

	GetUserNode(userID string) (*models.User, error)

	ListRoomMessages(
		roomID string,
		viewerID string,
		page int,
		pageSize int,
	) ([]*pb.MessageItem, error)

	ListMembers(
		context context.Context,
		request *pb.ListMembersRequest,
	) ([]*user.UserItem, error)

	SendMessage(
		messageID string,
		senderID string,
		roomId string,
		text string,
		replyToMessageId string,
	) (*pb.Message, error)
}

type chatRepository struct {
	cache   redisdb.RoomCache
	pubsub  *pubsub.Client
	spanner *gorm.DB
}

func NewChatRepository() ChatRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		DB:       0,
	})
	cache := redisdb.NewCache(
		time.Duration(time.Duration.Hours(1)),
		client,
	)
	pubsub, err := pubsub.NewClient(context.Background(), "cheers-a275e")
	if err != nil {
		pubsub = nil
	}

	return &chatRepository{
		cache:   cache,
		pubsub:  pubsub,
		spanner: utils.GetSpanner(),
	}
}
