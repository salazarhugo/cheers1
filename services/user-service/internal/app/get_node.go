package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
)

func (s *Server) GetUserNode(
	ctx context.Context,
	request *pb.GetUserNodeRequest,
) (*pb.GetUserNodeResponse, error) {
	userId := request.GetUserId()

	user, err := s.userRepository.GetUserNode(userId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.GetUserNodeResponse{
		User: user,
	}, nil
}
