package app

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error {
	_, err := utils.GetUserId(server.Context())
	if err != nil {
		return status.Error(codes.Internal, "failed to retrieve userID")
	}

	//err = s.chatRepository.JoinRoom(request, server)
	//if err != nil {
	//	return err
	//}
	//
	//return nil

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		DB:       0,
	})
	ctx := context.Background()

	sub := client.Subscribe(ctx, request.GetRoomId())
	defer sub.Close()
	log.Println(sub.String())

	message := pb.Message{}

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Println(err)
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", message)

		err = server.Send(&message)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("sent with no errors")
	}
}
