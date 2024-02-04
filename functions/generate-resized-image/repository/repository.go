package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func GetPost(url string) (*postpb.Post, error) {
	ctx := context.Background()

	conn := utils.CreateServiceConnection(ctx, "post-service-r3a2dr4u4a-nw.a.run.app:443")
	defer conn.Close()

	client := post.NewPostServiceClient(conn)

	response, err := client.GetPost(ctx, &post.GetPostRequest{PostId: url})
	if err != nil {
		log.Println(err)
	}
	return response.GetPost(), nil
}
