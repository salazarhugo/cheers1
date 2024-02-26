package app

import (
	"cloud.google.com/go/spanner"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePost(
	ctx context.Context,
	request *pb.CreatePostRequest,
) (*pb.PostResponse, error) {
	viewerID, err := utils.GetUserId(ctx)

	audio := request.GetAudio()
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving viewerID")
	}

	newPost := &models.Post{
		UserID:    viewerID,
		Caption:   request.Caption,
		Location:  request.LocationName,
		DrinkID:   spanner.NullString{Valid: false},
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	}

	if audio != nil {
		newPost.AudioUrl = audio.GetUrl()
		newPost.AudioWaveform = audio.GetWaveform()
	}

	if request.DrinkId != "" {
		newPost.DrinkID = spanner.NullString{StringVal: request.GetDrinkId(), Valid: true}
	}

	postID, err := s.postRepository.CreatePost(viewerID, newPost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create post")
	}

	err = s.postRepository.CreatePostMedias(postID, request.GetMediaIds())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to associate medias with post")
	}

	post, err := s.postRepository.GetPostItem(viewerID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get post")
	}

	err = pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
		Event: &pb.PostEvent_Create{
			Create: &pb.CreatePost{
				Post:                      post.GetPost(),
				User:                      nil,
				SendNotificationToFriends: true,
			},
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to publish event")
	}

	return post, nil
}
