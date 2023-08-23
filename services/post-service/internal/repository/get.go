package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (p *postRepository) GetPost(
	postID string,
) (*postpb.Post, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetPost.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"postID": postID,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	post := &postpb.Post{}

	if result.Next() {
		m := result.Record().Values[0]
		err := mapper.MapToProto(post, m)
		if err != nil {
			return nil, err
		}
	}

	return post, nil
}
