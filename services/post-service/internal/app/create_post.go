package app

import (
	"cloud.google.com/go/spanner"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePost(
	ctx context.Context,
	request *pb.CreatePostRequest,
) (*pb.PostResponse, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving viewerID")
	}

	newPost := &repository.Post{
		UserID:        viewerID,
		Caption:       request.Caption,
		Location:      request.LocationName,
		Photos:        request.Photos,
		DrinkID:       spanner.NullInt64{Valid: false},
		AudioUrl:      request.Audio.Url,
		AudioWaveform: request.Audio.Waveform,
	}
	if request.DrinkId > 0 {
		newPost.DrinkID = spanner.NullInt64{Int64: request.GetDrinkId(), Valid: true}
	}

	postID, err := s.postRepository.CreatePost(viewerID, newPost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create post")
	}

	post, err := s.postRepository.GetPostItem(viewerID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get post")
	}

	return post, nil
}
