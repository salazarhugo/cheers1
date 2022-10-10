package app

import (
	"context"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MicroserviceServer) CreatePost(
	ctx context.Context,
	request *v1.CreatePostRequest,
) (*post.Post, error) {
	_, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyReq := request.GetPost()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "post parameter can't be nil")
	}

	err = m.partyRepository.CreatePost(*partyReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return partyReq, nil
}
