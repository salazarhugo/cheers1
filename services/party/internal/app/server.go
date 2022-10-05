package app

import (
	"context"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"sync"
)

type MicroserviceServer struct {
	v1.UnimplementedMainServer
	mu              sync.Mutex
	partyRepository repository.PartyRepository
}

func NewMicroserviceServer() *MicroserviceServer {
	return &MicroserviceServer{
		v1.UnimplementedMainServer{},
		sync.Mutex{},
		repository.NewPartyRepository(utils.GetSession(utils.GetDriver())),
	}
}

func GetUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.InvalidArgument, "Failed retrieving metadata")
	}
	return md["user-id"][0], nil
}
