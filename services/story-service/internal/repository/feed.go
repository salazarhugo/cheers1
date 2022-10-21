package repository

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *storyRepository) FeedStory(
	userID string,
	request *pb.FeedStoryRequest,
) (*pb.FeedStoryResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	pageSize := request.GetPageSize()
	page := request.GetPage()

	skip := page * pageSize

	cypher, err := utils.GetCypher("internal/queries/FeedStory.cql")
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

	// Empty Feed
	storys := make([]*pb.StoryResponse, 0)

	for result.Next() {
		m := result.Record().Values[0]
		bytes, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		story := &pb.StoryResponse{}
		err = protojson.Unmarshal(bytes, story)
		if err != nil {
			return nil, err
		}
		storys = append(storys, story)
	}

	nextPageToken := "storyToken123"
	if len(storys) > 0 {
		nextPageToken = storys[len(storys)-1].Story.Id
	}

	return &pb.FeedStoryResponse{
		Stories:       storys,
		NextPageToken: nextPageToken,
	}, nil
}
