package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (p *userRepository) VerifyUser(
	userID string,
	verified bool,
) error {
	db := p.spanner

	err := db.
		Table("users").
		Where("UserId = ?", userID).
		Updates(models.User{Verified: verified}).
		Error
	if err != nil {
		return err
	}

	result, err := p.GetUserById(userID)
	if err != nil {
		return err
	}

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &pb.UserEvent{
		Event: &pb.UserEvent_Update{
			Update: &pb.UpdateUser{
				User: result.ToUserPb(),
			},
		},
	})

	return nil
}
