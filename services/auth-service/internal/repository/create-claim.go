package repository

import (
	"context"
	"fmt"
	claimgen "github.com/salazarhugo/cheers1/gen/go/cheers/claim/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/constants"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (a authRepository) CreateClaim(
	userID string,
	claim constants.Claim,
) error {
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

	claims[string(claim)] = true

	err = client.SetCustomUserClaims(ctx, userID, claims)
	if err != nil {
		log.Println(err)
		return status.Error(codes.Internal, fmt.Sprintf("error setting custom claim: %s", claim))
	}

	err = pubsub.PublishProtoWithBinaryEncoding("claim-topic", &claimgen.ClaimEvent{
		Event: &claimgen.ClaimEvent_Created{
			Created: &claimgen.CreatedClaim{
				UserId: userID,
				Claim:  string(claim),
			},
		},
	})

	return nil
}
