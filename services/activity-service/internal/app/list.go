package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListActivity(
	ctx context.Context,
	request *activity.ListActivityRequest,
) (*activity.ListActivityResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	activities, err := s.activityRepository.ListActivity(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list activity")
	}

	return &activity.ListActivityResponse{
		Activities: activities,
	}, nil
}
