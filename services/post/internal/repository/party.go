package repository

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/party"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type PartyRepository interface {
	CreateParty(party party.Party) error
	GetParty(id string) (*party.Party, error)
	UpdateParty(party party.Party) error
	DeleteParty(id string) error
	ListParty() ([]party.Party, error)
}

type partyRepository struct {
	session neo4j.Session
}

func NewPartyRepository(session neo4j.Session) PartyRepository {
	return &partyRepository{session: session}
}

func (p *partyRepository) CreateParty(party party.Party) error {
	cypher := `MATCH (u:User { id: $userid})
				CREATE (e: Party {
					id:            $id,
					name:          $name,
					description:   $description,
					address:       $address,
					latitude:      $latitude,
					longitude:     $longitude,
					privacy:       $privacy,
					banner_url:    $banner_url,
					start_date:    $start_date,
					end_date:      $end_date,
					host_id:       $host_id,
					location_name: $location_name,
					create_time:   $create_time
				})
		   		SET e += { created: datetime().epochMillis }
				CREATE (u)-[:POSTED]->(e)
				CREATE (u)-[:GOING]->(e)`

	params := map[string]interface{}{
		"userid":        party.HostId,
		"id":            uuid.NewString(),
		"name":          party.Name,
		"description":   party.Description,
		"address":       party.Address,
		"latitude":      party.GetLatlng().Latitude,
		"longitude":     party.GetLatlng().Longitude,
		"privacy":       party.Privacy.String(),
		"banner_url":    party.BannerUrl,
		"start_date":    party.StartDate.Seconds,
		"end_date":      party.EndDate.Seconds,
		"host_id":       party.HostId,
		"location_name": party.LocationName,
		"create_time":   timestamppb.Now().Seconds,
	}

	_, err := p.session.Run(cypher, params)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (p *partyRepository) GetParty(id string) (*party.Party, error) {
	cypher := `
			MATCH (u:User {id: $userId})
			MATCH (host:User)-[:POSTED]->(party: Party {id: $partyId})
			OPTIONAL MATCH (:User)-[interest:INTERESTED]->(party)
			OPTIONAL MATCH (:User)-[going:GOING]->(party)
			WITH 
				host,
				party, 
				exists((u)-[:INTERESTED]->(party)) as interested,
				exists((u)-[:GOING]->(party)) as going,
				count(DISTINCT interest) as interestedCount,
				count(DISTINCT going) as goingCount,
				host.id = u.id as isHost
			RETURN 
				party {
					.*,
					isHost: isHost,
					hostId: host.id,
					hostName: host.name,
					interested: interested,
					interestedCount: interestedCount,
					going: going,
					goingCount: goingCount
				}`

	params := map[string]interface{}{
		"userId":  "",
		"partyId": id,
	}

	result, err := p.session.Run(cypher, params)

	if err != nil {
		return nil, err
	}

	if result.Next() {
		a := result.Record().Values[0]
		log.Println(a)
	}

	return &party.Party{}, nil
}

func (p *partyRepository) UpdateParty(party party.Party) error {
	return nil
}

func (p *partyRepository) DeleteParty(id string) error {
	return nil
}

func (p *partyRepository) ListParty() ([]party.Party, error) {
	return []party.Party{}, nil
}
