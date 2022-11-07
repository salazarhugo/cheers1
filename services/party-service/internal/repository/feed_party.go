package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *partyRepository) FeedParty(
	userID string,
	request *pb.FeedPartyRequest,
) (*pb.FeedPartyResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	pageSize := 50
	skip := 0

	cypher, err := utils.GetCypher("internal/queries/FeedParty.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":   userID,
		"skip":     int(skip),
		"pageSize": int(pageSize),
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	items := make([]*pb.PartyItem, 0)

	for result.Next() {
		party := &pb.PartyItem{}
		data := result.Record().Values[0]
		err = utils.MapToProto(party, data)
		if err != nil {
			return nil, err
		}
		items = append(items, party)
	}

	return &pb.FeedPartyResponse{
		Items:         items,
		NextPageToken: "token123",
	}, nil
}
