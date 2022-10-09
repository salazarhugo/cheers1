package app

import (
	"context"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MicroserviceServer) ListPost(
	ctx context.Context,
	request *v1.ListPostRequest,
) (*v1.ListPostResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	posts, err := m.partyRepository.ListPost(userID, request.GetPageToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return posts, nil
}
