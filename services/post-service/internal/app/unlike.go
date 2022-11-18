package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UnlikePost(
	ctx context.Context,
	request *pb.UnlikePostRequest,
) (*pb.UnlikePostResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	response, err := s.postRepository.UnlikePost(userID, request.PostId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
