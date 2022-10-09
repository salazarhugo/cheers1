package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MicroserviceServer) DeleteParty(
	ctx context.Context,
	request *v1.DeletePartyRequest,
) (*empty.Empty, error) {
	_, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyID := request.GetId()
	if partyID == "" {
		return nil, status.Error(codes.InvalidArgument, "partyID parameter can't be empty")
	}

	err = m.partyRepository.DeleteParty(partyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return &empty.Empty{}, nil
}
