package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDrink(
	ctx context.Context,
	request *drink.CreateDrinkRequest,
) (*drink.CreateDrinkResponse, error) {
	_, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	if request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty field: name")
	}

	drinkID, err := s.repository.CreateDrink(
		request.Name,
		request.Icon,
	)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	res, err := s.repository.GetDrink(drinkID)

	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &drink.CreateDrinkResponse{
		Drink: res,
	}, nil
}
