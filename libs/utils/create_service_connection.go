package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func CreateServiceConnection(
	ctx context.Context,
	target string,
) *grpc.ClientConn {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	conn, err := grpc.DialContext(ctx, target,
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithUnaryInterceptor(CloudRunInterceptor),
	)

	return conn
}
