package utils

import (
	"context"
	"fmt"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"time"
)

func CloudRunInterceptor(
	// The incoming request received from the gateway or another service.
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}

	audience, err := getAudience(cc.Target())
	if err != nil {
		log.Println(err)
	}

	accessToken, err := generateAccessToken(ctx, audience)
	if err != nil {
		log.Println(err)
	}

	log.Println(audience, accessToken)

	md.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	md.Set("X-Serverless-Authorization", fmt.Sprintf("Bearer %s", accessToken))

	ctx = metadata.NewOutgoingContext(ctx, md)

	// Calls the invoker to execute RPC
	err = invoker(ctx, method, req, reply, cc, opts...)

	// Logic after invoking the invoker
	log.Printf("Invoked RPC method=%s; Duration=%s; Error=%v", method, time.Since(start), err)

	return err
}

func getAudience(target string) (string, error) {
	host, _, err := net.SplitHostPort(target)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s", host), nil
}

func generateAccessToken(ctx context.Context, audience string) (string, error) {
	tokenSource, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		return "", err
	}

	token, err := tokenSource.Token()
	if err != nil {
		log.Println(err)
		return "", nil
	}

	return token.AccessToken, nil
}
