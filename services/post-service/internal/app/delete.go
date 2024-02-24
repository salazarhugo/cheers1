package app

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeletePost(
	ctx context.Context,
	request *pb.DeletePostRequest,
) (*emptypb.Empty, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving viewerID")
	}

	postID := request.GetId()
	if postID == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be empty")
	}

	post, err := s.postRepository.GetPostItem(viewerID, postID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "post not found")
	}

	if post.User.Id != viewerID {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("user %s is not authorized to delete this post", viewerID))
	}

	err = s.postRepository.DeletePost(postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete post")
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Delete{
				Delete: &pb.DeletePost{
					Post: post.GetPost(),
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return &empty.Empty{}, nil
}
