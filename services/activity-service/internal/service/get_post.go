package service

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func GetPost(postId string) (*postpb.Post, error) {
	ctx := context.Background()
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	conn, err := grpc.DialContext(ctx, "post-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := post.NewPostServiceClient(conn)

	response, err := client.GetPost(ctx, &post.GetPostRequest{PostId: postId})
	if err != nil {
		log.Println(err)
	}
	return response.GetPost(), nil
}
