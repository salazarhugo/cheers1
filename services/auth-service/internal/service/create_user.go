package service

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func CreateUser(
	email string,
	username string,
) (*user.User, error) {
	ctx := context.Background()
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	conn, err := grpc.DialContext(ctx, "user-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithUnaryInterceptor(utils.CloudRunInterceptor),
	)
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	request := &userpb.CreateUserRequest{
		Email:    email,
		Username: username,
	}

	response, err := client.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.GetUser(), nil
}
