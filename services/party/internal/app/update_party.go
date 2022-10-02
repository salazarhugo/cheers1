package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"log"
)

func (s *MainApiServiceServer) UpdateParty(
	ctx context.Context,
	request *v1.CreatePartyRequest,
) (*empty.Empty, error) {
	log.Println(request)
	log.Println("UpdateParty!")

	return &empty.Empty{}, nil
}
