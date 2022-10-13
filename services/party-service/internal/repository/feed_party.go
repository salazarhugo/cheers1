package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/party/v1"
)

func (p *partyRepository) FeedParty(
	userID string,
	request *pb.FeedPartyRequest,
) (*pb.FeedPartyResponse, error) {

	parties := make([]*pb.PartyResponse, 0)

	return &pb.FeedPartyResponse{
		Parties:       parties,
		NextPageToken: "token123",
	}, nil
}
