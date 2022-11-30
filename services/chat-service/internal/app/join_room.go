package app

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (s *Server) JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error {
	//_, err := utils.GetUserId(server.Context())
	//if err != nil {
	//	return status.Error(codes.Internal, "failed to retrieve userID")
	//}

	err := s.chatRepository.JoinRoom(request, server)
	if err != nil {
		return err
	}

	return err
}
