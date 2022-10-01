package main

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	proto "github.com/salazarhugo/cheers1/genproto"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"log"
	"sync"
)

func NewServer() *MainApiServiceServer {
	s := &MainApiServiceServer{
		proto.UnimplementedMainServer{},
		sync.Mutex{},
		utils.GetSession(utils.GetDriver()),
	}
	fmt.Println(s)
	return s
}

type MainApiServiceServer struct {
	proto.UnimplementedMainServer
	mu sync.Mutex
	neo4j.Session
}

type Party struct {
	Id           string  `json:"id" structs:"id"`
	Name         string  `json:"name" structs:"name"`
	Description  string  `json:"description" structs:"description"`
	Address      string  `json:"address" structs:"address"`
	Privacy      string  `json:"privacy" structs:"privacy"`
	BannerUrl    string  `json:"bannerUrl" structs:"bannerUrl"`
	StartDate    int64   `json:"startDate" structs:"startDate"`
	EndDate      int64   `json:"endDate" structs:"endDate"`
	HostId       string  `json:"hostId" structs:"hostId"`
	Longitude    float64 `json:"longitude" structs:"longitude"`
	Latitude     float64 `json:"latitude" structs:"latitude"`
	LocationName string  `json:"locationName" structs:"locationName"`
}

func (s *MainApiServiceServer) CreateParty(
	ctx context.Context,
	request *proto.CreatePartyRequest,
) (*proto.CreatePartyResponse, error) {
	log.Println(request)

	cypher := `MATCH (u:User { username: $username}) 
				CREATE (e: Party $party)
		   		SET e += { created: datetime().epochMillis } 
				CREATE (u)-[:POSTED]->(e)
				CREATE (u)-[:GOING]->(e)`

	params := map[string]interface{}{
		"username": "hugosalazar",
		"party": structs.Map(Party{
			Id:           "awfwafwaf",
			Name:         "Goland",
			Description:  "description",
			Address:      "",
			Privacy:      "FRIENDS",
			BannerUrl:    "",
			StartDate:    0,
			EndDate:      0,
			HostId:       "omg",
			Longitude:    3,
			Latitude:     3,
			LocationName: "Germasoyeia",
		}),
	}

	_, err := s.Session.Run(cypher, params)
	if err != nil {
		log.Println(err)
	}

	return &proto.CreatePartyResponse{}, nil
}

func (s *MainApiServiceServer) DeleteParty(
	ctx context.Context,
	request *proto.DeletePartyRequest,
) (*empty.Empty, error) {
	log.Println(request)
	return &empty.Empty{}, nil
}
