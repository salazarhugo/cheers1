package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *postRepository) LikePost(
	userID string,
	postID string,
) (*pb.LikePostResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/LikePost.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID": userID,
		"postID": postID,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	go func() {
		postResponse, err := p.GetPost(userID, postID)
		post := postResponse.GetPost()
		err = pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Like{
				Like: &pb.LikePost{
					Post: post,
					User: nil,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return &pb.LikePostResponse{
		Success: true,
	}, nil
}
