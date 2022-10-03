package app

import (
	"context"
	"github.com/fatih/structs"
	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (s *MainApiServiceServer) CreateParty(
	ctx context.Context,
	request *v1.CreatePartyRequest,
) (*party.Party, error) {
	log.Println(request)
	log.Println(ctx.Value("userId"))
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Failed retrieving metadata")
	}
	log.Println(md)

	partyReq := request.GetParty()

	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "party parameter cant be nil")
	}

	log.Println(partyReq)

	cypher := `MATCH (u:User { username: $username})
				CREATE (e: Party $party)
		   		SET e += { created: datetime().epochMillis }
				CREATE (u)-[:POSTED]->(e)
				CREATE (u)-[:GOING]->(e)`

	params := map[string]interface{}{
		"username": "hugosalazar",
		"party": structs.Map(party.Party{
			Id:           "goland",
			Name:         partyReq.Name,
			Description:  partyReq.Description,
			Address:      partyReq.Address,
			Latlng:       partyReq.Latlng,
			Privacy:      partyReq.Privacy,
			BannerUrl:    partyReq.BannerUrl,
			StartDate:    partyReq.StartDate,
			EndDate:      partyReq.EndDate,
			HostId:       "goland",
			LocationName: partyReq.LocationName,
			CreateTime:   timestamppb.Now(),
		}),
	}

	_, err := s.Session.Run(cypher, params)
	if err != nil {
		log.Println(err)
	}

	return partyReq, nil
}
