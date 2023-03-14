package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListSuggestions(
	ctx context.Context,
	request *pb.ListSuggestionsRequest,
) (*pb.ListSuggestionsResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	suggestions, err := s.userRepository.ListSuggestions(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.ListSuggestionsResponse{
		Users: suggestions,
	}, nil
}
