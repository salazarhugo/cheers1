package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/location/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListFriendLocation(
	ctx context.Context,
	request *location.ListFriendLocationRequest,
) (*location.ListFriendLocationResponse, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve viewerID")
	}

	items, err := s.repository.ListFriendLocation(
		ctx,
		viewerID,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &location.ListFriendLocationResponse{
		Items: items,
	}, nil
}
