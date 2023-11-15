package repository

import (
	uuid "github.com/google/uuid"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *postRepository) CreatePost(
	userID string,
	post *Post,
) (string, error) {
	db := p.spanner

	post.ID = uuid.NewString()
	post.UserID = userID

	result := db.Create(&post)
	if result.Error != nil {
		return "", result.Error
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Create{
				Create: &pb.CreatePost{
					Post: &postpb.Post{
						Id:                    post.ID,
						CreatorId:             userID,
						Caption:               post.Caption,
						Address:               "",
						Privacy:               0,
						Photos:                nil,
						LocationName:          "",
						Drink:                 "",
						Drunkenness:           0,
						Type:                  0,
						CreateTime:            0,
						CanComment:            false,
						CanShare:              false,
						Ratio:                 0,
						Latitude:              0,
						Longitude:             0,
						LastCommentText:       "",
						LastCommentUsername:   "",
						LastCommentCreateTime: 0,
					},
					User:                      nil,
					SendNotificationToFriends: true,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return post.ID, nil
}
