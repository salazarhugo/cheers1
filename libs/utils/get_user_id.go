package utils

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.InvalidArgument, "failed retrieving metadata")
	}

	if len(md["user-id"]) < 1 {
		return "", status.Error(codes.InvalidArgument, "missing user-id")
	}

	return md["user-id"][0], nil
}
