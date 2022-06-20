package main

import (
	"chat/cache"
	"chat/chatpb"
	"chat/neo4j"
	"context"
	json2 "encoding/json"
	"errors"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"sync"
	"time"
)

var (
	roomCache cache.RoomCache
)

func newServer() *chatServiceServer {
	s := &chatServiceServer{
		channel: make(map[string][]chan *chatpb.Message),
	}
	fmt.Println(s)
	roomCache = cache.NewRedisCache(
		"redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		"mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		0,
		time.Duration(time.Duration.Hours(1)),
	)
	return s
}

type chatServiceServer struct {
	chatpb.UnimplementedChatServiceServer
	mu      sync.Mutex
	channel map[string][]chan *chatpb.Message
}

func (s *chatServiceServer) GetRooms(_ *chatpb.Empty, roomStream chatpb.ChatService_GetRoomsServer) error {
	md, err := authorize(roomStream.Context())
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	rooms := roomCache.GetRooms(userId)

	for i := range rooms {
		if err := roomStream.Send(rooms[i]); err != nil {
			log.Printf("send error %v", err)
		}
	}

	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		DB:       0,
	})
	sub := client.Subscribe(ctx, "messages")
	defer sub.Close()
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		message := chatpb.Message{}
		if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
			panic(err)
		}

		room := roomCache.GetRoomWithId(userId, message.GetRoom().GetId())

		if err := roomStream.Send(room); err != nil {
			log.Printf("send error %v", err)
		}
	}
}

func (s *chatServiceServer) CreateChat(ctx context.Context, request *chatpb.CreateChatReq) (*chatpb.Room, error) {
	md, err := authorize(ctx)
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	UUIDs := append([]string{userId}, request.UserIds...)
	UUIDs = unique(UUIDs)

	if len(UUIDs) < 2 {
		return nil, errors.New("chats must have at least 2 users")
	}

	// Direct Chat
	if len(UUIDs) == 2 {
		room := roomCache.GetOrCreateDirectRoom(UUIDs[0], UUIDs[1])
		return room, nil
	}

	room := roomCache.CreateGroup(request.GroupName, UUIDs)
	return room, nil
}

func (s *chatServiceServer) LeaveRoom(ctx context.Context, request *chatpb.RoomId) (*chatpb.Empty, error) {
	md, err := authorize(ctx)
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	log.Println("Leave Room with id: ", request.GetRoomId())
	roomCache.LeaveRoom(userId, request.GetRoomId())

	return &chatpb.Empty{}, nil
}

func (s *chatServiceServer) GetRoomMembers(ctx context.Context, request *chatpb.RoomId) (*chatpb.Users, error) {
	_, err := authorize(ctx)
	if err != nil {
		panic(err)
	}

	membersId := roomCache.GetRoomMembers(request.RoomId)
	users, err := neo4j.GetUsers(membersId)
	if err != nil {
		return nil, err
	}

	return &chatpb.Users{Users: users}, nil
}

func (s *chatServiceServer) DeleteRoom(ctx context.Context, request *chatpb.RoomId) (*chatpb.Empty, error) {
	_, err := authorize(ctx)
	if err != nil {
		panic(err)
	}

	roomCache.DeleteRoom(request.GetRoomId())

	return &chatpb.Empty{}, nil
}

func (s *chatServiceServer) GetRoomId(ctx context.Context, request *chatpb.GetRoomIdReq) (*chatpb.RoomId, error) {
	md, err := authorize(ctx)
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	room := roomCache.GetOrCreateDirectRoom(userId, request.GetRecipientId())

	return &chatpb.RoomId{RoomId: room.GetId()}, nil
}

func (s *chatServiceServer) JoinRoom(request *chatpb.RoomId, msgStream chatpb.ChatService_JoinRoomServer) error {
	md, err := authorize(msgStream.Context())
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	var ctx = context.Background()

	isMember := roomCache.IsMember(userId, request.RoomId)

	if isMember == false {
		log.Println("Can't access rooms you are not member of")
		return errors.New("you are not member of this group chat")
	}

	go roomCache.SetSeen(request.GetRoomId(), userId)

	messages := roomCache.GetMessages(request.GetRoomId())

	for _, msg := range messages {
		if err := msgStream.Send(msg); err != nil {
			log.Printf("send error %v", err)
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		DB:       0,
	})
	sub := client.Subscribe(ctx, request.GetRoomId())
	defer sub.Close()
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		message := chatpb.Message{}
		if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
			panic(err)
		}
		log.Println(message)
		err = msgStream.Send(&message)
		if err != nil {
			return err
		}
	}
}

func (s *chatServiceServer) SendMessage(msgStream chatpb.ChatService_SendMessageServer) error {
	fmt.Println("Send Message...")
	md, err := authorize(msgStream.Context())
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		DB:       0,
	})

	msg, err := msgStream.Recv()

	isMember := roomCache.IsMember(userId, msg.Room.Id)
	if isMember == false {
		log.Println("Can't access rooms you are not member of")
		return errors.New("you are not member of this group chat")
	}

	go SendMessageNotification(userId, msg.Room.Id, msg.SenderName, msg.SenderProfilePictureUrl)
	go roomCache.SetSeen(msg.Room.Id, userId)

	roomCache.SetMessage(msg)

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	json, err := json2.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = client.Publish(msgStream.Context(), msg.GetRoom().GetId(), json).Err()
	err = client.Publish(msgStream.Context(), "messages", json).Err()
	if err != nil {
		log.Println("could not publish to channel", err)
	}

	ack := chatpb.MessageAck{Status: "SENT"}
	msgStream.SendAndClose(&ack)

	go func() {
		streams := s.channel[msg.Room.GetId()]
		for _, msgChan := range streams {
			msgChan <- msg
		}
	}()

	return nil
}

func (s *chatServiceServer) DeleteUser(ctx context.Context, req *chatpb.UserIdReq) (*chatpb.Empty, error) {
	DeleteUser(req.GetUserId())
	return &chatpb.Empty{}, nil
}

func (s *chatServiceServer) LikeMessage(ctx context.Context, req *chatpb.LikeMessageReq) (*chatpb.Empty, error) {
	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Failed retrieving metadata")
	}
	//userId := md.Get("userId")[0]

	roomCache.LikeMessage(req.RoomId, req.MessageId)

	return &chatpb.Empty{}, nil
}

func (s *chatServiceServer) AddToken(ctx context.Context, req *chatpb.AddTokenReq) (*chatpb.Empty, error) {
	md, err := authorize(ctx)
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]
	roomCache.AddToken(userId, req.Token)

	return &chatpb.Empty{}, err
}

func (s *chatServiceServer) TypingStart(ctx context.Context, req *chatpb.TypingReq) (*chatpb.Empty, error) {
	md, err := authorize(ctx)
	if err != nil {
		panic(err)
	}
	userId := md.Get("userId")[0]

	go SendNotification(userId, req.RoomId, req.Username, req.AvatarUrl)
	return &chatpb.Empty{}, nil
}

func SendMessageNotification(userId string, roomId string, name string, avatar string) {
	app := initializeAppDefault()
	a, err := app.Messaging(context.Background())
	if err != nil {
		return
	}

	members := roomCache.GetRoomMembers(roomId)
	var tokens []string

	for _, uid := range members {
		if uid != userId {
			tokens = append(tokens, roomCache.GetTokens(uid)...)
		}
	}

	_, err = a.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Data: map[string]string{
			"type":      "newMessage",
			"title":     name,
			"body":      "sent a Chat",
			"avatar":    avatar,
			"roomId":    roomId,
			"channelId": "CHAT_CHANNEL",
		},
		Tokens: tokens,
	})

	if err != nil {
		log.Println(err)
	}
}

func SendNotification(userId string, roomId string, name string, avatar string) {
	app := initializeAppDefault()
	a, err := app.Messaging(context.Background())
	if err != nil {
		return
	}

	members := roomCache.GetRoomMembers(roomId)
	var tokens []string

	for _, uid := range members {
		if uid != userId {
			tokens = append(tokens, roomCache.GetTokens(uid)...)
		}
	}

	_, err = a.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Data: map[string]string{
			"type":      "newMessage",
			"title":     name,
			"body":      "is typing...",
			"avatar":    avatar,
			"roomId":    roomId,
			"channelId": "CHAT_CHANNEL",
		},
		Tokens: tokens,
	})

	if err != nil {
		log.Println(err)
	}
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
