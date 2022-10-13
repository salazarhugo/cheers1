package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

func GetCypher(name string) (*string, error) {
	bytes, err := os.ReadFile(name)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed reading cql file")
	}

	cypher := string(bytes)

	return &cypher, nil
}
