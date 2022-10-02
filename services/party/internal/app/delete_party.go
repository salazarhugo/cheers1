package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"log"
)

func (s *MainApiServiceServer) DeleteParty(
	ctx context.Context,
	request *v1.DeletePartyRequest,
) (*empty.Empty, error) {
	log.Println(request)
	return &empty.Empty{}, nil
}
