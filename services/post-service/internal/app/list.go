package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) FeedPost(ctx context.Context, request *pb.FeedPostRequest) (*pb.FeedPostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	posts, err := s.partyRepository.FeedPost(userID, request)
	if err != nil {
		return nil, err
	}

	return posts, nil
}