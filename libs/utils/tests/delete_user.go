package tests

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"gorm.io/gorm"
)

func DeleteUser(
	db *gorm.DB,
	userID string,
	publishEvent bool,
) error {
	result := db.Delete(&models.User{UserId: userID})

	if result.Error != nil {
		return result.Error
	}

	if publishEvent {
		err := pubsub.PublishProtoWithBinaryEncoding("user-topic", &pb.UserEvent{
			Event: &pb.UserEvent_Delete{
				Delete: &pb.DeleteUser{
					UserId: userID,
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
