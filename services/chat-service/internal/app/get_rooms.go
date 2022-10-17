package app

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRooms(
	empty *pb.Empty,
	server pb.ChatService_GetRoomsServer,
) error {
	_, err := utils.GetUserId(server.Context())
	if err != nil {
		return status.Error(codes.Internal, "Failed retrieving userID")
	}

	//rooms := s.chatRepository.GetRooms(userID)
	//
	//for i := range rooms {
	//	if err := server.Send(rooms[i]); err != nil {
	//		log.Printf("send error %v", err)
	//	}
	//}
	//
	//var ctx = context.Background()
	//
	//sub := s.chatRepository.Subscribe(ctx, "messages")
	//defer sub.Close()
	//for {
	//	msg, err := sub.ReceiveMessage(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//	message := chatpb.Message{}
	//	if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
	//		panic(err)
	//	}
	//
	//	room := roomCache.GetRoomWithId(userId, message.GetRoom().GetId())
	//
	//	if err := roomStream.Send(room); err != nil {
	//		log.Printf("send error %v", err)
	//	}
	//}
	return nil
}
