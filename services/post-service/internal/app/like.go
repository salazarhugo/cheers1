package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LikePost(
	ctx context.Context,
	request *pb.LikePostRequest,
) (*pb.LikePostResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	response, err := s.postRepository.LikePost(userID, request.Id)
	if err != nil {
		return nil, err
	}

	return response, nil
}
