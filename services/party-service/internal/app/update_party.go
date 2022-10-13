package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/party/v1"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
)

func (s *Server) UpdateParty(
	ctx context.Context,
	request *pb.UpdatePartyRequest,
) (*party.Party, error) {
	return &party.Party{}, nil
}
