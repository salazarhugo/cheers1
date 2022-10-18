package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/party/v1"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
)

type PartyRepository interface {
	CreateParty(userID string, party *party.Party) (*party.Party, error)
	GetParty(id string) (*party.Party, error)
	UpdateParty(party *party.Party) error
	DeleteParty(id string) error

	FeedParty(userID string, request *pb.FeedPartyRequest) (*pb.FeedPartyResponse, error)
}

type partyRepository struct {
	driver neo4j.Driver
}

func NewPartyRepository(driver neo4j.Driver) PartyRepository {
	return &partyRepository{
		driver: driver,
	}
}
