package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *postRepository) DeletePost(
	userID string,
	postID string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/DeletePost.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"postID": postID,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	go func() {
		//postResponse, err := p.GetPost(userID, postID)
		//post := postResponse.GetPost()
		err = pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Delete{
				Delete: &pb.DeletePost{
					Sender: nil,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
