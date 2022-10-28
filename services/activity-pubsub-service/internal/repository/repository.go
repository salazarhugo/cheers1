package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

type ActivityRepository interface {
	CreateActivity(event *post.PostEvent) error
}

type activityRepository struct {
}

func (a activityRepository) CreateActivity(event *post.PostEvent) error {
	ctx := context.Background()
	client, err := utils.InitializeAppDefault().Firestore(ctx)
	if err != nil {
		return err
	}

	_, err = client.Collection("activities").Doc(userID).Create(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository() ActivityRepository {
	return &activityRepository{}
}

func GetPost(postID string) (*post.Post, error) {
	ctx := context.Background()

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	conn, err := grpc.DialContext(ctx, os.Getenv("POST_SERVICE_HOST")+":443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := post.NewPostServiceClient(conn)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("not ok")
	}
	log.Println(md)
	ctx := metadata.AppendToOutgoingContext(ctx, "authorization", md.Get("x-forwarded-authorization")[0])

	response, err := client.GetPost(ctx, &post.GetPostRequest{Id: postID})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}
