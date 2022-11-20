package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) FeedPost(
	userID string,
	request *pb.FeedPostRequest,
) (*pb.FeedPostResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	pageSize := request.GetPageSize()
	page := request.GetPage()

	skip := page * pageSize

	cypher, err := utils.GetCypher("internal/queries/FeedPost.cql")
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
	posts := make([]*pb.PostResponse, 0)

	for result.Next() {
		m := result.Record().Values[0]
		post := &pb.PostResponse{}
		err := utils.MapToProto(post, m)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	nextPageToken := "postToken123"
	if len(posts) > 0 {
		nextPageToken = posts[len(posts)-1].Post.Id
	}

	return &pb.FeedPostResponse{
		Posts:         posts,
		NextPageToken: nextPageToken,
	}, nil
}
