package app

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeletePost(
	ctx context.Context,
	request *pb.DeletePostRequest,
) (*emptypb.Empty, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	postID := request.GetId()
	if postID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be empty")
	}

	post, err := s.postRepository.GetPost(userID, postID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "post not found")
	}

	if post.Post.CreatorId != userID {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user %s is not authorized to delete this post", userID))
	}

	err = s.postRepository.DeletePost(userID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete post")
	}

	return &empty.Empty{}, nil
}
