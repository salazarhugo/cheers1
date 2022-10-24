package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteStory(
	ctx context.Context,
	request *pb.DeleteStoryRequest,
) (*emptypb.Empty, error) {
	_, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	storyID := request.GetId()
	if storyID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be blank")
	}

	err = s.storyRepository.DeleteStory(storyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete story")
	}

	return &emptypb.Empty{}, err
}
