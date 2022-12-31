package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListMapPost(
	ctx context.Context, request *pb.ListMapPostRequest,
) (*pb.ListMapPostResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	posts, err := s.postRepository.ListMapPost(userID, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return posts, nil
}
