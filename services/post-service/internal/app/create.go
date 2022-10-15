package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePost(
	ctx context.Context,
	request *pb.CreatePostRequest,
) (*pb.PostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyReq := request.GetPost()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "post parameter can't be nil")
	}

	postID, err := s.partyRepository.CreatePost(userID, partyReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create post")
	}

	post, err := s.partyRepository.GetPost(userID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get post")
	}

	return post, nil
}
