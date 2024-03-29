package repository

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	chatmodels "github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/redisdb"
	"gorm.io/gorm"
	"os"
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

	GetInbox(
		userID string,
	) ([]*pb.RoomWithMessages, error)

	GetChat(
		userID string,
		chatID string,
	) (*chatmodels.Chat, error)

	ListChatsIds(
		userID string,
	) ([]string, error)

	DeleteRoom(
		userID string,
		roomID string,
	) error

	DeleteRooms(
		userID string,
	) error

	GetUserNode(
		userID string,
	) (*models.User, error)

	GetUsersNode(
		userIDs []string,
	) ([]*models.User, error)

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
		chatID string,
		text string,
		replyToMessageId string,
	) (*chatmodels.ChatMessage, error)

	SendReadReceipt(
		userID string,
		chatID string,
	) error
}

type chatRepository struct {
	cache   redisdb.RoomCache
	pubsub  *pubsub.Client
	spanner *gorm.DB
}

func (c chatRepository) GetChat(userID string, chatID string) (*chatmodels.Chat, error) {
	return c.cache.GetChat(userID, chatID)
}

func (c chatRepository) ListChatsIds(userID string) ([]string, error) {
	return c.cache.ListChatIds(userID)
}

func (c chatRepository) SendReadReceipt(userID string, chatID string) error {
	err := c.cache.SetLastRead(chatID, userID, time.Now().UnixMilli())
	if err != nil {
		return err
	}

	return c.cache.ResetUnreadCount(chatID, userID)
}

func NewChatRepository(
	redisClient *redis.Client,
) ChatRepository {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("DB_ENDPOINT"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       0,
		})
	}

	cache := redisdb.NewCache(
		time.Duration(time.Duration.Hours(1)),
		redisClient,
	)

	pubsub, err := pubsub.NewClient(
		context.Background(),
		"cheers-a275e",
	)
	if err != nil {
		pubsub = nil
	}

	return &chatRepository{
		cache:   cache,
		pubsub:  pubsub,
		spanner: utils.GetSpanner(),
	}
}
