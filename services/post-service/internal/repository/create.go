package repository

import (
	uuid "github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
	"time"
)

func (p *postRepository) CreatePost(
	userID string,
	request *pb.CreatePostRequest,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreatePost.cql")
	if err != nil {
		return "", err
	}

	post := request.GetPost()
	post.Id = uuid.NewString()
	post.CreatorId = userID
	post.CreateTime = time.Now().Unix()

	m, err := mapper.ProtoToMap(post)
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
		err := pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Create{
				Create: &pb.CreatePost{
					Post:                      post,
					User:                      nil,
					SendNotificationToFriends: request.GetSendNotificationToFriends(),
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return post.Id, nil
}
