package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePost(
	ctx context.Context,
	request *pb.CreatePostRequest,
) (*pb.PostResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyReq := request.GetPost()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "post parameter can't be nil")
	}
	if partyReq.Caption == "" && len(partyReq.Photos) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "empty caption and photos")
	}

	postReq := &repository.Post{
		Caption: partyReq.Caption,
	}

	postID, err := s.postRepository.CreatePost(userID, postReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create post")
	}

	post, err := s.postRepository.GetPostItem(userID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get post")
	}

	return post, nil
}
