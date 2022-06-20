package main

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

func streamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	md, err := authorize(ss.Context())
	if err != nil {
		return err
	}

	err = ss.SetHeader(md)
	ss.SendHeader(md)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Set headers stream", md)

	return handler(srv, ss)
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	if info.FullMethod == "/chatpb.ChatService/DeleteUser" {
		log.Println("Skipped authentification")
		h, err := handler(ctx, req)
		return h, err
	}

	start := time.Now()

	md, err := authorize(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	h, err := handler(ctx, req)

	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return md, status.Errorf(codes.InvalidArgument, "Failed retrieving metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return md, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}
	token := authHeader[0]

	// validateToken function validates the token
	validToken, err := validateToken(token)
	if err != nil {
		return md, status.Errorf(codes.Unauthenticated, err.Error())
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "userId", validToken.UID)

	return metadata.Pairs("userId", validToken.UID), nil
}

func validateToken(idToken string) (*auth.Token, error) {
	app := initializeAppDefault()
	client := initializeAuthClient(app)

	idToken = strings.Split(idToken, "Bearer ")[1]

	token, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		log.Println("Invalid Authorization Token", idToken)
		return nil, err
	}
	return token, nil
}
