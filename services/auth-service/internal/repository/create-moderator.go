package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/claim/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (a authRepository) CreateModerator(userID string) error {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return err
	}

	otherUser, err := client.GetUser(ctx, userID)

	claims := otherUser.CustomClaims
	if claims == nil {
		claims = make(map[string]interface{}, 0)
	}

	claims["moderator"] = true

	err = client.SetCustomUserClaims(ctx, userID, claims)
	if err != nil {
		log.Println(err)
		return status.Error(codes.Internal, "error setting custom claims")
	}

	err = pubsub.PublishProtoWithBinaryEncoding("claim-topic", &claim.ClaimEvent{
		Event: &claim.ClaimEvent_Created{
			Created: &claim.CreatedClaim{
				UserId: userID,
				Claim:  "moderator",
			},
		},
	})

	return nil
}
