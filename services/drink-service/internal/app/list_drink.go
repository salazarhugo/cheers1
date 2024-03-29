package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListDrink(
	ctx context.Context,
	request *drink.ListDrinkRequest,
) (*drink.ListDrinkResponse, error) {
	_, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	items, err := s.repository.ListDrink()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &drink.ListDrinkResponse{
		Items: items,
	}, nil
}
