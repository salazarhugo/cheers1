package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/services/postservice/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListPost(ctx context.Context, request *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	posts, err := s.partyRepository.ListPost(userID, request.GetPageToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return posts, nil
}
