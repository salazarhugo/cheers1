package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
	"gorm.io/gorm"
)

type PartyRepository interface {
	CreateParty(userID string, party *model.Party) (string, error)
	GetPartyById(id string) (*model.Party, error)
	GetPartyBySlug(slug string) (*model.Party, error)
	GetPartyWithName(name string) (*party.Party, error)
	UpdateParty(party *model.Party) (string, error)
	UpsertPartyBySlug(party *model.Party) (string, error)
	DeletePartyById(id string) error

	TransferParty(userID string, partyID string) error
	GetPartyItem(userID string, partyID string) (*pb.PartyItem, error)
	FeedParty(userID string, request *pb.FeedPartyRequest) (*pb.FeedPartyResponse, error)
	ListParty(viewerID string, request *pb.ListPartyRequest) (*pb.ListPartyResponse, error)
	UpdateWatchStatus(userID string, partyID string, status pb.WatchStatus) error
	ListGoing(userID string, partyID string) ([]*user.UserItem, error)
	UnGoingParty(userID string, partyID string) error
	UpdateMinimumPrice(price int64, partyID string) error
}

type partyRepository struct {
	driver  neo4j.Driver
	spanner *gorm.DB
}

func NewRepository() PartyRepository {
	return &partyRepository{
		driver:  utils.GetDriver(),
		spanner: utils.GetSpanner(),
	}
}
