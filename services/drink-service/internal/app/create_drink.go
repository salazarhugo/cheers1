package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateDrink(
	ctx context.Context,
	request *drink.CreateDrinkRequest,
) (*drink.CreateDrinkResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.repository.CreateDrink(
		userID,
		request.Name,
		request.Icon,
		request.Category,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &drink.CreateDrinkResponse{}, nil
}
