package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"sync"
)

type Server struct {
	party.UnimplementedPartyServiceServer
	mu              sync.Mutex
	partyRepository repository.PartyRepository
}

func NewServer() *Server {
	return &Server{
		party.UnimplementedPartyServiceServer{},
		sync.Mutex{},
		repository.NewPartyRepository(utils.GetDriver()),
	}
}
