package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) UpdateDrink(
	ctx context.Context,
	request *drink.UpdateDrinkRequest,
) (*drink.UpdateDrinkResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.repository.UpdateDrink(
		userID,
		request.Latitude,
		request.Longitude,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &drink.UpdateDrinkResponse{}, nil
}
