package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/salazarhugo/cheers1/proto/out/cheers/api/v1"
	"log"
	"sync"
)

func NewServer() *MainApiServiceServer {
	s := &MainApiServiceServer{}
	fmt.Println(s)
	return s
}

type MainApiServiceServer struct {
	v1.UnimplementedMainServer
	mu sync.Mutex
}

func (s *MainApiServiceServer) CreateParty(
	ctx context.Context,
	request *v1.CreatePartyRequest,
) (*v1.CreatePartyResponse, error) {
	log.Println(request)
	request.
	return &v1.CreatePartyResponse{}, nil
}

func (s *MainApiServiceServer) DeleteParty(
	ctx context.Context,
	request *v1.DeletePartyRequest,
) (*empty.Empty, error) {
	log.Println(request)
	return &empty.Empty{}, nil
}
