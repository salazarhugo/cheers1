package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
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
		err := utils.PublishProtoMessages("post-topic", &pb.PostEvent{
			UserId: userID,
			PostId: postID,
			Type:   pb.PostEvent_DELETE,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
