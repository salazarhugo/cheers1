package app

import (
	"context"
	"github.com/fatih/structs"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type"
	"log"
)

func (s *MainApiServiceServer) CreateParty(
	ctx context.Context,
	request *v1.CreatePartyRequest,
) (*party.Party, error) {
	log.Println(request)

	cypher := `MATCH (u:User { username: $username}) 
				CREATE (e: Party $party)
		   		SET e += { created: datetime().epochMillis } 
				CREATE (u)-[:POSTED]->(e)
				CREATE (u)-[:GOING]->(e)`

	params := map[string]interface{}{
		"username": "hugosalazar",
		"party": structs.Map(party.Party{
			Id:   "awfwafwaf",
			Name: "Goland",
		}),
	}

	_, err := s.Session.Run(cypher, params)
	if err != nil {
		log.Println(err)
	}

	return &party.Party{
		Id:   "party-01",
		Name: "Guaba Saturday",
	}, nil
}
