package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/claim/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateBusinessAccount(
	ctx context.Context,
	request *auth.CreateBusinessAccountRequest,
) (*auth.CreateBusinessAccountResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	// Check if user is admin
	user, err := client.GetUser(ctx, userID)
	if user.CustomClaims["admin"] == nil {
		return nil, status.Error(codes.PermissionDenied, "insufficient permissions")
	}

	otherUser, err := client.GetUser(ctx, request.UserId)

	claims := otherUser.CustomClaims
	if claims == nil {
		claims = make(map[string]interface{}, 0)
	}

	claims["business"] = true

	err = client.SetCustomUserClaims(ctx, request.UserId, claims)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "error setting custom claims")
	}

	err = pubsub.PublishProtoWithBinaryEncoding("claim-topic", &claim.ClaimEvent{
		Event: &claim.ClaimEvent_Created{
			Created: &claim.CreatedClaim{
				UserId: request.UserId,
				Claim:  "business",
			},
		},
	})

	return &auth.CreateBusinessAccountResponse{}, nil
}
