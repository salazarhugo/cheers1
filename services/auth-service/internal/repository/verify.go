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

func (r authRepository) VerifyUser(
	userID string,
) error {
	ctx := context.Background()

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return err
	}

	otherUser, err := client.GetUser(ctx, userID)
	if err != nil {
		return status.Error(codes.Internal, "failed to get auth user")
	}

	claims := otherUser.CustomClaims
	if claims == nil {
		claims = make(map[string]interface{}, 0)
	}

	claims["verified"] = true

	err = client.SetCustomUserClaims(ctx, userID, claims)
	if err != nil {
		log.Println(err)
		return status.Error(codes.Internal, "error setting custom claims")
	}

	err = pubsub.PublishProtoWithBinaryEncoding("claim-topic", &claim.ClaimEvent{
		Event: &claim.ClaimEvent_Created{
			Created: &claim.CreatedClaim{
				UserId: userID,
				Claim:  "verified",
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
