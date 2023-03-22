package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetDrink(
	ctx context.Context,
	request *drink.GetDrinkRequest,
) (*drink.GetDrinkResponse, error) {
	_, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	res, err := s.repository.GetDrink(
		request.GetDrinkId(),
	)

	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &drink.GetDrinkResponse{
		Drink: res,
	}, nil
}
