package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

type PartyRepository interface {
	CreateParty(userID string, party *party.Party) (*party.Party, error)
	GetParty(id string) (*party.Party, error)
	UpdateParty(party *party.Party) (string, error)
	DeleteParty(id string) error

	TransferParty(userID string, partyID string) error
	GetPartyItem(userID string, partyID string) (*pb.PartyItem, error)
	GetPartyItemPublic(partyID string) (*pb.PartyItem, error)
	FeedParty(userID string, request *pb.FeedPartyRequest) (*pb.FeedPartyResponse, error)
	ListParty(viewerID string, request *pb.ListPartyRequest) (*pb.ListPartyResponse, error)
	UpdateWatchStatus(userID string, partyID string, status pb.WatchStatus) error
	ListGoing(userID string, partyID string) ([]*user.UserItem, error)
	UnGoingParty(userID string, partyID string) error
	UpdateMinimumPrice(price int64, partyID string) error
}

type partyRepository struct {
	driver neo4j.Driver
}

func NewPartyRepository(driver neo4j.Driver) PartyRepository {
	return &partyRepository{
		driver: driver,
	}
}
