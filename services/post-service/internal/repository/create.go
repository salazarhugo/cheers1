package repository

import (
	uuid "github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"log"
	"time"
)

func (p *postRepository) CreatePost(
	userID string,
	post *postpb.Post,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreatePost.cql")
	if err != nil {
		return "", err
	}

	post.Id = uuid.NewString()
	post.CreatorId = userID
	post.CreateTime = int32(time.Now().Unix())

	m, err := utils.ProtoToMap(post)
	if err != nil {
		return "", err
	}

	params := map[string]interface{}{
		"userID": userID,
		"post":   m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return "", err
	}

	go func() {
		err := utils.PublishProtoMessages("post-topic", &pb.PostEvent{
			UserId:       userID,
			PostId:       post.Id,
			CreatorId:    post.CreatorId,
			Caption:      post.Caption,
			Address:      post.Address,
			LocationName: post.LocationName,
			Drink:        post.Drink,
			Type:         pb.PostEvent_CREATE,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return post.Id, nil
}
