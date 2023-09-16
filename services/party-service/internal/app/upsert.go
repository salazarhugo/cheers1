package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpsertParty(
	ctx context.Context,
	request *pb.UpsertPartyRequest,
) (*pb.UpsertPartyResponse, error) {
	partyReq := request.GetParty()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "party parameter can't be nil")
	}

	party, err := s.partyRepository.GetPartyWithSlug(partyReq.Slug)
	if party != nil {
		partyReq.Id = party.Id
		newPartyId, err := s.partyRepository.UpdateParty(partyReq)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to update party")
		}

		newParty, err := s.partyRepository.GetParty(newPartyId)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to get new party")
		}

		return &pb.UpsertPartyResponse{
			Party: newParty,
		}, nil
	}

	newPartyId, err := s.partyRepository.InsertParty(partyReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to insert party")
	}

	newParty, err := s.partyRepository.GetParty(newPartyId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get new party")
	}

	return &pb.UpsertPartyResponse{
		Party: newParty,
	}, nil
}
