package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*emptypb.Empty, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	postID := request.GetId()
	if postID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be empty")
	}

	log.Print(postID)

	err = s.postRepository.DeletePost(userID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return &empty.Empty{}, nil
}
